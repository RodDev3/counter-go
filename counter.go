package main

// Counter struct
type Counter struct {
	Attempts int `json:"attempts"`
	Success int `json:"success"`
	SuccessPerfect int `json:"successPerfect"`
	CurrentStreak int `json:"currentStreak"`
	BestStreak int `json:"bestStreak"`
}

// NewApp creates a new App application struct
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) IncrementSuccess() Counter {
	c.Success++
	return *c
}

func (c *Counter) DecrementSuccess() Counter {
	if c.Success > 0 {
		c.Success--
	}
	return *c
}

func (c *Counter) IncrementAttempts() Counter {
	c.Attempts++
	return *c
}

func (c *Counter) DecrementAttempts() Counter {
	if c.Attempts > 0 {
		c.Attempts--
	}
	return *c
}

func (c *Counter) IncrementSuccessPerfect() Counter {
	c.SuccessPerfect++
	return *c
}


func (c *Counter) DecrementSuccessPerfect() Counter {
	if c.SuccessPerfect > 0 {
		c.SuccessPerfect--
	}
	return *c
}

func (c *Counter) IncrementCurrentStreak() Counter {
	c.CurrentStreak++
	return *c
}


func (c *Counter) DecrementCurrentStreak() Counter {
	if c.CurrentStreak > 0 {
		c.CurrentStreak--
	}
	return *c
}

func (c *Counter) IncrementBestStreak() Counter {
	c.BestStreak++
	return *c
}


func (c *Counter) DecrementBestStreak() Counter {
	if c.BestStreak > 0 {
		c.BestStreak--
	}
	return *c
}

func (c *Counter) GetCounter() Counter {
	return *c
}

func (c *Counter) AutoSuccess() Counter {
	c.Attempts++
	c.Success++
	c.CurrentStreak++
	if c.CurrentStreak > c.BestStreak {
		c.BestStreak = c.CurrentStreak
	}

	return *c
}

func (c *Counter) AutoSuccessPerfect() Counter {
	c.SuccessPerfect++
	c.AutoSuccess()

	return *c
}

func (c *Counter) AutoFail() Counter {
	c.Attempts++
	c.CurrentStreak = 0

	return *c
}