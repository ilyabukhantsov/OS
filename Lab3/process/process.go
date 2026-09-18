package process

import(
"Lab3/pageTableEntries"
	)

type Process struct{
			ID int
			Table []pageTableEntries.PageTableEntries
			Saved []bool
}
