// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package depstests

import (
	"testing"
	"time"

	"github.com/golang/protobuf/jsonpb" //nolint:staticcheck
	"github.com/golang/protobuf/ptypes" //nolint:staticcheck
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDuration(t *testing.T) {
	// https://github.com/golang/protobuf/issues/883
	// https://github.com/golang/protobuf/issues/1219
	// https://jira.percona.com/browse/PMM-6760

	var m jsonpb.Marshaler
	s, err := m.MarshalToString(ptypes.DurationProto(-time.Nanosecond))
	require.NoError(t, err)
	assert.Equal(t, `"-0.000000001s"`, s)
}
