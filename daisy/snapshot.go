//  Copyright 2020 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package daisy

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	daisyCompute "github.com/GoogleCloudPlatform/compute-image-tools/daisy/compute"
	"google.golang.org/api/compute/v1"
)

var (
	snapshotURLRgx = regexp.MustCompile(fmt.Sprintf(`^(projects/(?P<project>%[1]s)/)?global/snapshots/(?P<snapshot>%[2]s)$`, projectRgxStr, rfc1035))
)

// snapshotExists should only be used during validation for existing GCE snapshots
// and should not be relied or populated for daisy created resources.
func (w *Workflow) snapshotExists(project, snapshot string) (bool, DError) {
	return w.snapshotCache.resourceExists(func(project string, opts ...daisyCompute.ListCallOption) (interface{}, error) {
		return w.ComputeClient.ListSnapshots(project)
	}, project, snapshot)
}

// Snapshot is used to create a GCE disk in a project.
type Snapshot struct {
	compute.Snapshot
	Resource

	sourceDiskProject string
	sourceDiskZone    string
	sourceDiskName    string
}

// MarshalJSON is a hacky workaround to prevent Snapshot from using compute.Snapshot's implementation.
func (ss *Snapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(*ss)
}

func (ss *Snapshot) populate(ctx context.Context, s *Step) DError {
	var errs DError
	ss.Name, errs = ss.Resource.populateWithGlobal(ctx, s, ss.Name)

	ss.Description = strOr(ss.Description, fmt.Sprintf("Snapshot created by Daisy in workflow %q on behalf of %s.", s.w.Name, s.w.username))

	if diskURLRgx.MatchString(ss.SourceDisk) {
		ss.SourceDisk = extendPartialURL(ss.SourceDisk, ss.Project)
	}

	m := NamedSubexp(diskURLRgx, ss.SourceDisk)
	ss.sourceDiskProject = m["project"]
	ss.sourceDiskZone = m["zone"]
	ss.sourceDiskName = m["disk"]

	ss.link = fmt.Sprintf("projects/%s/global/snapshots/%s", ss.Project, ss.Name)
	return errs
}

func (ss *Snapshot) validate(ctx context.Context, s *Step) DError {
	pre := fmt.Sprintf("cannot create snapshot %q", ss.daisyName)
	errs := ss.Resource.validate(ctx, s, pre)

	// Source disk checking.
	if ss.SourceDisk == "" {
		errs = addErrs(errs, Errf("%s: must provide SourceDisk", pre))
	}
	if _, err := s.w.disks.regUse(ss.SourceDisk, s); err != nil {
		errs = addErrs(errs, newErr("failed to get source disk", err))
	}

	// Register creation.
	errs = addErrs(errs, s.w.snapshots.regCreate(ss.daisyName, &ss.Resource, s, false))
	return errs
}

type snapshotRegistry struct {
	baseResourceRegistry
}

func newSnapshotRegistry(w *Workflow) *snapshotRegistry {
	sr := &snapshotRegistry{baseResourceRegistry: baseResourceRegistry{w: w, typeName: "snapshot", urlRgx: snapshotURLRgx}}
	sr.baseResourceRegistry.deleteFn = sr.deleteFn
	sr.init()
	return sr
}