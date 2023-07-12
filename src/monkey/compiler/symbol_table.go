package compiler



type SymbolScope string

const (
	GlobalScope   SymbolScope = "GLOBAL"
	LocalScope    SymbolScope = "LOCAL"
	BuiltinScope  SymbolScope = "BUILTIN"
	FreeScope     SymbolScope = "FREE"
    // the name of the function we’re currently compiling.
	FunctionScope SymbolScope = "FUNCTION"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int // symbol operand, used for setGlobal, getGlobal
}

// recursive SymbolTable, similar to  recursive environment in interpreter
type SymbolTable struct {
	Outer *SymbolTable
	// symbol name to symbol mapping
	store          map[string]Symbol
	numDefinitions int
	FreeSymbols    []Symbol
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	free := []Symbol{}
	return &SymbolTable{store: s, FreeSymbols: free}
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

// binding an Identifier to symbol table
func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions}
	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}
	// fmt.Printf("Define %s in scope %s\n", name, symbol.Scope)
	s.store[name] = symbol
	s.numDefinitions++

	return symbol
}

//	func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
//		obj, ok := s.store[name]
//		if !ok && s.Outer != nil {
//			obj, ok = s.Outer.Resolve(name)
//			return obj, ok
//		}
//		return obj, ok
//	}
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	// resolve as locals(parameters + locals)
	obj, ok := s.store[name]
	if !ok && s.Outer != nil {
		// resolve as outers, could be free variable, global variable or buitins
		obj, ok = s.Outer.Resolve(name)
		// if it's free variable, then it has to be in outer scope
		if !ok {
			return obj, ok
		}
		if obj.Scope == GlobalScope || obj.Scope == BuiltinScope {
			return obj, ok
		}
		// noteworthy: free variables are defined when being resolved
		// resolve as free variables
		free := s.defineFree(obj)
		return free, true
	}
	return obj, ok
}

func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{Name: name, Index: index, Scope: BuiltinScope}
	s.store[name] = symbol
	return symbol
}

func (s *SymbolTable) defineFree(original Symbol) Symbol {
	// original is a local in outer scope
	s.FreeSymbols = append(s.FreeSymbols, original)
	// save original as free in inner scope
	symbol := Symbol{Name: original.Name, Index: len(s.FreeSymbols) - 1}
	symbol.Scope = FreeScope
	s.store[original.Name] = symbol
	return symbol
}

func (s *SymbolTable) DefineFunctionName(name string) Symbol {
	symbol := Symbol{Name: name, Index: 0, Scope: FunctionScope}
	s.store[name] = symbol
	return symbol
}
