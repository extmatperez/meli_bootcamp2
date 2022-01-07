/*  Ejercicio 1 - Service/Repo/Db Update()
Diseñar un test que pruebe en la capa service, el método o función Update(). Para lograrlo se deberá:
Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado.
Para dar el test como OK debe validarse que al invocar el método del Service Update(),  retorne el
producto con mismo Id y los datos actualizados. Validar también que  Read() del Repository haya sido
ejecutado durante el test.


 Ejercicio 2 - Service/Repo/Db Delete()
Diseñar un test que pruebe en la capa service, el método o función Delete(). Se debe probar la correcta
eliminación de un producto, y el error cuando el producto no existe. Para lograrlo puede:
Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
Ejecutar el test con dos id’s de producto distintos, siendo uno de ellos un id inexistente en el Mock de
Storage.
Para dar el test como OK debe validarse que efectivamente el producto borrado ya no exista en Storage
luego del Delete(). También que cuando se intenta borrar un producto  inexistente, se debe obtener el
error correspondiente.

*/

package main
