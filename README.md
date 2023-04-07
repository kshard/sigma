# sigma

algebra of σ-calculus

## Quick Example

```go
rules := ast.Rules{/* rules of σ-expression */}

machine, err := sigma.New(/* goal */, rules)
if err != nil {
	panic(err)
}

ctx := asm.NewContext().Add(/* add generators */)
reader := sigma.Stream(ctx, machine)

reader.ToSeq()
```

## Internal

* `ast` - abstract syntax tree to define the σ-calculus expression
* `internal/complier` - compiles the σ-calculus expression into assembler ("byte code")
* `asm` - assembler of the VM
* `vm` - Sigma VM

```go
// Compile phase
rules := ast.Rules{/* rules of σ-expression */}

build := compiler.New()
build.Compile(rules)

machine, shape, obj := c.Assemble(/* goal */)

// Execution phase
ctx := asm.NewContext().Add(/* add generators */)

stream := obj.Link(ctx)
reader := machine.Stream(shape, stream)

reader.ToSeq()
```