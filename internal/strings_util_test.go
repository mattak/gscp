package internal

import "testing"

func TestSplitBucketURI(t *testing.T) {
	tf := TestFrame{Test: t}

	tf.Run("gs://", func(t TestFrame) {
		name, obj := SplitBucketURI("gs://")
		tf.AssertEquals(name, "")
		tf.AssertEquals(obj, "")
	})

	tf.Run("gs://sample", func(t TestFrame) {
		name, obj := SplitBucketURI("gs://sample")
		tf.AssertEquals(name, "sample")
		tf.AssertEquals(obj, "")
	})

	{
		name, obj := SplitBucketURI("gs://sample/path1")
		tf.AssertEquals(name, "sample")
		tf.AssertEquals(obj, "path1")
	}

	{
		name, obj := SplitBucketURI("gs://sample/path1/")
		tf.AssertEquals(name, "sample")
		tf.AssertEquals(obj, "path1/")
	}

	{
		name, obj := SplitBucketURI("gs://sample/path1/path2")
		tf.AssertEquals(name, "sample")
		tf.AssertEquals(obj, "path1/path2")
	}
}
