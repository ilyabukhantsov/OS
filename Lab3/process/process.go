package process

import(
"Lab3/pageTableEntries"
	)

type Process struct{
			ID int
			Table []pageTableEntries.PageTableEntries
			WorkingSet []int
}

func newProcess(ID int, Table []pageTableEntries.PageTableEntries, WorkingSet []int) *Process{
	return &Process{
		ID: ID,
		Table: Table,
		WorkingSet: WorkingSet,
	}
}

func (*Process) Work() error{
	return nil
}
