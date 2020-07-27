# APUNTES Learn go with test(Aplicando TDD)


## Discipline
- Let's go over the cycle again
- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

## The TDD process and why the steps are important
- Write a failing test and see it fail as first step
- Write a relevant test for our requirements 
- Produces an easy to understand description of the failure(escenario)
- Writing the smallest amount of code to make it pass so we know we have working software(baby steps)
- Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work wit

## Break Down

- Escribir el código en pequeñas funciones testeables
- Funciones lo más genéricas posibles que se puedan reutilizar

## Write the minimal amount of code for the test to run and check the failing test output
- Keep the discipline! You don't need to know anything new right now to make the test fail properly.
- All you need to do right now is enough to make it compile so you can check your test is written well

## Hacerlo un hábito
- Quizás no se pueda aplicar en todas las situaciones dado un proyecto, sin embargo, aplicarlo
a consciencia lo más repetidamente posible ayudará a incorporarlo de forma más innata.
- Obligarse a escribir el código mínimo para compilar es muy importante.

## Expecting behaviors
- Escribimos funcionalidades esperando que se comporten de cierta forma, lo comprobamos con los test unitarios.
- Cómo se comporta o qué lógica sigue la aplicación para manejar distintos escenarios.
- Qué pasa si a esta función le doy como parámetro x, que responde, cómo comunica los errores, etc.

## Benchmarking
- Otra feature de go
- Se escriben similares a los test
- Comando go test -bench=.
- Ej se va a repetir b.N veces, esto lo decide el framework internamente

```golang
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```
## Benchmarking output
- Se repitió la función 7489444 veces para determinar que el promedio de ejecución promedio es 155 nanosegundos

```s
goos: linux
goarch: amd64
pkg: iteration
BenchmarkRepeat-8        7489444               155 ns/op
PASS
ok      iteration       1.330s

```