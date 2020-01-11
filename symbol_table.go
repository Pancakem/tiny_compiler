package main

const MAX_SYMBOL_TABLE_SIZE = 100

type TableItem struct {
	id   int
	val  int
	name string
}

var (
	symbol_table [MAX_SYMBOL_TABLE_SIZE]TableItem
	table_size   = 0
)

func addSym(name string) int {

	for i := 0; i < table_size; i++ {
		if symbol_table[i].name == name {
			return symbol_table[i].id
		}

	}

	item := TableItem{}
	item.id = table_size
	item.name = name

	symbol_table[table_size] = item
	table_size += 1
	return table_size
}

func getSym(id int) *TableItem {
	return &symbol_table[id]
}

func setSym(id, val int) {
	symbol_table[id].val = val
}
