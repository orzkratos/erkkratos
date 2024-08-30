package erkrequire

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/stretchr/testify/require"
)

func No(t require.TestingT, erk *errors.Error, msgAndArgs ...interface{}) {
	require.NoError(t, eke(erk), msgAndArgs...)
}

func Eo(t require.TestingT, erk *errors.Error, msgAndArgs ...interface{}) {
	require.Error(t, eke(erk), msgAndArgs...) //这里必须传递个空才行，跟前面的情况相同
}

func eke(erk *errors.Error) error {
	if erk == nil {
		return nil //这里必须做这样的转换，因为两个 nil 是不一样的
	}
	return erk
}

func Is(t require.TestingT, expected *errors.Error, actualEx *errors.Error, msgAndArgs ...interface{}) {
	require.Equal(t, bool(eke(expected) == nil), bool(eke(actualEx) == nil), msgAndArgs...)
	if expected != nil && actualEx != nil {
		require.Equal(t, expected.Reason, actualEx.Reason, msgAndArgs...)
		require.Equal(t, expected.Code, actualEx.Code, msgAndArgs...)
	}
}
