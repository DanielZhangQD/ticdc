package processor

import (
	"testing"

	"github.com/pingcap/check"
)

func TestSuite(t *testing.T) {
	check.TestingT(t)
}

type etcdkeySuite struct{}

var _ = check.Suite(&etcdkeySuite{})

func (s *etcdkeySuite) TestEtcdKey(c *check.C) {
	testcases := []struct {
		key      string
		expected *CDCEtcdKey
	}{{
		key: "/tidb/cdc/owner/223176cb44d20a13",
		expected: &CDCEtcdKey{
			Tp:           CDCEtcdKeyTypeOnwer,
			OwnerLeaseID: "223176cb44d20a13",
		},
	}, {
		key: "/tidb/cdc/owner",
		expected: &CDCEtcdKey{
			Tp:           CDCEtcdKeyTypeOnwer,
			OwnerLeaseID: "",
		},
	}, {
		key: "/tidb/cdc/capture/6bbc01c8-0605-4f86-a0f9-b3119109b225",
		expected: &CDCEtcdKey{
			Tp:        CDCEtcdKeyTypeCapture,
			CaptureID: "6bbc01c8-0605-4f86-a0f9-b3119109b225",
		},
	}, {
		key: "/tidb/cdc/changefeed/info/test-_@#$%changefeed",
		expected: &CDCEtcdKey{
			Tp:           CDCEtcdKeyTypeChangefeedInfo,
			ChangefeedID: "test-_@#$%changefeed",
		},
	}, {
		key: "/tidb/cdc/changefeed/info/test/changefeed",
		expected: &CDCEtcdKey{
			Tp:           CDCEtcdKeyTypeChangefeedInfo,
			ChangefeedID: "test/changefeed",
		},
	}, {
		key: "/tidb/cdc/job/test-changefeed",
		expected: &CDCEtcdKey{
			Tp:           CDCEtcdKeyTypeChangeFeedStatus,
			ChangefeedID: "test-changefeed",
		},
	}, {
		key: "/tidb/cdc/task/position/6bbc01c8-0605-4f86-a0f9-b3119109b225/test-changefeed",
		expected: &CDCEtcdKey{
			Tp:           CDCEtcdKeyTypeTaskPosition,
			ChangefeedID: "test-changefeed",
			CaptureID:    "6bbc01c8-0605-4f86-a0f9-b3119109b225",
		},
	}, {
		key: "/tidb/cdc/task/position/6bbc01c8-0605-4f86-a0f9-b3119109b225/test/changefeed",
		expected: &CDCEtcdKey{
			Tp:           CDCEtcdKeyTypeTaskPosition,
			ChangefeedID: "test/changefeed",
			CaptureID:    "6bbc01c8-0605-4f86-a0f9-b3119109b225",
		},
	}, {
		key: "/tidb/cdc/task/status/6bbc01c8-0605-4f86-a0f9-b3119109b225/test-changefeed",
		expected: &CDCEtcdKey{
			Tp:           CDCEtcdKeyTypeTaskStatus,
			ChangefeedID: "test-changefeed",
			CaptureID:    "6bbc01c8-0605-4f86-a0f9-b3119109b225",
		},
	}, {
		key: "/tidb/cdc/task/workload/6bbc01c8-0605-4f86-a0f9-b3119109b225/test-changefeed",
		expected: &CDCEtcdKey{
			Tp:           CDCEtcdKeyTypeTaskWorkload,
			ChangefeedID: "test-changefeed",
			CaptureID:    "6bbc01c8-0605-4f86-a0f9-b3119109b225",
		},
	}}
	for _, tc := range testcases {
		k := new(CDCEtcdKey)
		err := k.Parse(tc.key)
		c.Assert(err, check.IsNil)
		c.Assert(k, check.DeepEquals, tc.expected)
		c.Assert(k.String(), check.Equals, tc.key)
	}
}

func (s *etcdkeySuite) TestEtcdKeyParseError(c *check.C) {
	testCases := []struct {
		key   string
		error bool
	}{{
		key:   "/tidb/cdc/task/position/6bbc01c8-0605-4f86-a0f9-b3119109b225/test/changefeed",
		error: false,
	}, {
		key:   "/tidb/cdc/task/position/6bbc01c8-0605-4f86-a0f9-b3119109b225/",
		error: false,
	}, {
		key:   "/tidb/cdc/task/position/6bbc01c8-0605-4f86-a0f9-b3119109b225",
		error: true,
	}, {
		key:   "/tidb/cdc/task/status/6bbc01c8-0605-4f86-a0f9-b3119109b225",
		error: true,
	}, {
		key:   "/tidb/cdc/task/workload/6bbc01c8-0605-4f86-a0f9-b3119109b225",
		error: true,
	}, {
		key:   "/tidb/cd",
		error: true,
	}, {
		key:   "/tidb/cdc/",
		error: true,
	}}
	for _, tc := range testCases {
		k := new(CDCEtcdKey)
		err := k.Parse(tc.key)
		if tc.error {
			c.Assert(err, check.NotNil)
		} else {
			c.Assert(err, check.IsNil)
		}
	}
}