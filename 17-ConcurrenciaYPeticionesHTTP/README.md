# Introducción

En programacion existen dos conceptos que regularmente son confundidos como si fueran la misma cosa, los cuales son *`concurrencia`* y *`paralelismo`*.

## Concurrencia

La concurrencia se define como el hecho de que un procesador tenga la capacidad de realizar multiples tareas intercaladamente, es decir, supongamos que tenemos 3 tareas las cuales demandan mucho tiempo para poder realizarse en un proceso computacional, si esto lo ejecutaramos de forma secuencial tendriamos un problema de rendimiento bastante grabe, ya que el procesador estaria ejecutando un tarea y hasta que esta tarea no termina ejecuta la siguiente y luego la siguiente y asi sucesivamente hasta ejecutar todas las tareas, esto tiene un problema y es que las tareas que quedan al final de la cola de procesos tardan mucho en ser realizadas. Es aqui donde entra la concurrencia ya que la concurrencia nos permite que el procesador vaya delegando las tareas mientras las ejecuta simultaneamente, es decir, mientras realiza una tarea si esta tarda mucho la deja ejecutandose y mientras sigue con otra tarea, despues esa otra tarea la deja ejecutandose y se va a otra tarea, despues regresa a la primera y continua con su ejecucion donde se quedo y asi sucesivamente.

## Paralelismo

El paralelismo por su parte consiste en tener varios procesadores ejecutando posiblemente la misma tarea al mismo tiempo, pero estos no intervienen en las tareas que ejecutan los demas procesadores, simplemente se encargan de realizar sus tareas y ya.