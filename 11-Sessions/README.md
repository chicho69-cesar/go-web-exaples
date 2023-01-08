# Introducción

Este ejemplo mostrará cómo almacenar datos en cookies de sesión usando el popular paquete `gorilla/sessions` en Go.
Las cookies son pequeños fragmentos de datos que se almacenan en el navegador de un usuario y se envían a nuestro servidor con cada solicitud. En ellos, podemos almacenar, por ejemplo, si un usuario ha iniciado sesión o no en nuestro sitio web y averiguar quién es realmente (en nuestro sistema).
En este ejemplo, solo permitiremos que los usuarios autenticados vean nuestro mensaje secreto en la página `/secret`. Para obtener acceso a él, primero tendrá que visitar `/login` para obtener una cookie de sesión válida, que lo registra. Además, puede visitar `/logout` para revocar su acceso a nuestro mensaje secreto.