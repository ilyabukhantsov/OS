package process

import(
"Lab3/pageTableEntries"
	)

type Process struct{
			ID int
			RemainingLife int
			Table []*pageTableEntries.PageTableEntries
			WorkingSet []int
}

func newProcess(ID int, RemainingLife int,  StartFrom int, numberOfPage int, WorkingSet []int) *Process{
	arrayOfPages := []*pageTableEntries.PageTableEntries{}
	for index := range numberOfPage{
		arrayOfPages = append(arrayOfPages, pageTableEntries.NewPageTableEntries(index + StartFrom))
	}
	return &Process{
		ID: ID,
		RemainingLife: RemainingLife,
		Table: arrayOfPages,
		WorkingSet: WorkingSet,
	}
}

func (*Process) Work() error{
	return nil
}
