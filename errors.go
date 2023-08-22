package ssh

type AllUnavailable struct{}

func (AllUnavailable) Error() string { return `all SSH client targets are unavailable` }
