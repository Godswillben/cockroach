// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package kvserverpb

import (
	"math"

	"github.com/cockroachdb/cockroach/pkg/storage/enginepb"
	"github.com/cockroachdb/cockroach/pkg/util/hlc"
)

var maxRaftCommandFooterSize = (&RaftCommandFooter{
	MaxLeaseIndex: math.MaxUint64,
	ClosedTimestamp: hlc.Timestamp{
		WallTime:  math.MaxInt64,
		Logical:   math.MaxInt32,
		Synthetic: true,
	},
}).Size()

// MaxRaftCommandFooterSize returns the maximum possible size of an encoded
// RaftCommandFooter proto.
func MaxRaftCommandFooterSize() int {
	return maxRaftCommandFooterSize
}

// IsZero returns whether all fields are set to their zero value.
func (r ReplicatedEvalResult) IsZero() bool {
	return r == ReplicatedEvalResult{}
}

// IsTrivial implements apply.Command.
//
// Trivial commands can be combined in an apply.Batch for efficiency. We only
// allow this for commands that don't have any side effects except a few updates
// that are part of all commands (MVCCStats, WriteTimestamp, etc), as other side
// effects may not be correctly propagated between commands in the same batch.
//
// User writes are generally trivial, so nontrivial commands are the exception,
// notable examples being configuration changes, splits, merges, etc.
//
// At the time of writing it is possible that the current conditions are too
// strict but they are certainly sufficient.
func (r *ReplicatedEvalResult) IsTrivial() bool {
	// Check if there are any non-trivial State updates.
	if r.State != nil {
		stateAllowlist := *r.State
		// ReplicaState.Stats was previously non-nullable which caused nodes to
		// send a zero-value MVCCStats structure. If the proposal was generated by
		// an old node, we'll have decoded that zero-value structure setting
		// ReplicaState.Stats to a non-nil value which would trigger the "unhandled
		// field in ReplicatedEvalResult" assertion to fire if we didn't clear it.
		// TODO(ajwerner): eliminate this case that likely can no longer occur as of
		// at least 19.1.
		if stateAllowlist.Stats != nil && (*stateAllowlist.Stats == enginepb.MVCCStats{}) {
			stateAllowlist.Stats = nil
		}
		if stateAllowlist != (ReplicaState{}) {
			return false
		}
	}
	// Set allowlist to the value of r and clear the allowlisted fields.
	// If allowlist is zero-valued after clearing the allowlisted fields then
	// it is trivial.
	allowlist := *r
	allowlist.Delta = enginepb.MVCCStatsDelta{}
	allowlist.WriteTimestamp = hlc.Timestamp{}
	allowlist.PrevLeaseProposal = nil
	allowlist.State = nil
	return allowlist.IsZero()
}
