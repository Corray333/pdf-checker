<template>
  <main>
    <div class="wrapper">
      <div class="header">
        <a href="/">
          <span class="logo-container">
          <img src="./assets/logo.png" alt="" class="logo">
          <div>
            <p>OsaSoft</p>
            <p>PDF Checker</p>
          </div>
        </span>
        </a>
        <a href="">Наши контакты</a>
      </div>
      <div class="content">
        <Transition>
          <label class="file-select" v-if="result == ''">
            <h1>Загрузите файл</h1>
            <!-- We can't use a normal button element here, as it would become the target of the label. -->
            <div class="select-button">
              <!-- Display the filename if a file has been selected. -->
              <span>+</span>
            </div>
            <!-- Now, the file input that we hide. -->
            <input id="fileInput" type="file" @change="uploadFile" />
          </label>
        </Transition>
        <Transition>
          <div v-if="result != ''" class="errors">
            <!-- <object :data="filePath" type="application/pdf" width="100%" height="400px" id="preview">
              <p>Unable to display PDF file. <a
                  href="/uploads/media/default/0001/01/540cb75550adf33f281f29132dddd14fded85bfc.pdf">Download</a> instead.
              </p>
            </object> -->
            <h2>Список ошибок:</h2>
            <ul>
              <li v-for="el in result">{{ el }}</li>
            </ul>
          </div>
        </Transition>
      </div>
    </div>
  </main>
</template>

<script setup>
import { ref } from "vue"

const result = ref("")

function uploadFile() {
  // Получаем элемент input
  let input = document.getElementById("fileInput")
  // Проверяем, что файл выбран
  if (input.files.length > 0) {
    // Создаем объект FormData
    let formData = new FormData()
    // Добавляем файл в объект FormData
    formData.append("myFile", input.files[0])
    // Создаем объект XMLHttpRequest
    let xhr = new XMLHttpRequest()
    // Открываем соединение с сервером
    xhr.open("POST", "/loadpdf")
    // Отправляем запрос с данными формы
    xhr.send(formData)
    // Обрабатываем ответ сервера
    xhr.onload = function () {
      console.log(xhr.responseText)
      result.value = JSON.parse(xhr.responseText)
    }
  } else {
    // Выводим сообщение, если файл не выбран
    alert("Пожалуйста, выберите файл для загрузки")
  }
}


</script>

<style scoped>

main {
  overflow: hidden;
  padding: 25px 350px;
  height: 100vh;
  margin: 0;
}

.wrapper {
  border: 4px solid var(--accent);
  border-radius: 30px;
  height: 100%;
  width: 100%;
}

.header {
  transform: scale(1.01) translateY(-1px);
  border: 4px solid var(--accent);
  background-color: var(--accent);
  border-radius: 30px;
  padding: 0 15px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.header>a {
  color: var(--black);
  text-decoration: none;
  transition: 0.2s ease-out;
}

.header>a:hover {
  transform: scale(1.1);
  transition: 0.2s ease-out;
}

.logo {
  height: 50px;
  border-radius: 25px;
}

.logo-container {
  display: flex;
  font-weight: bold;
  align-items: center;
  gap: 5px;
  padding: 5px 0;
}

.content {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

label {
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  gap: 15px;
}

.select-button {
  width: 100px;
  height: 100px;
  padding: 10px;
  color: white;
  background-color: var(--accent);
  font-size: 64px;
  border-radius: 50px;
  text-align: center;
  font-weight: bold;
  transition: 0.3s ease-out;
}

.select-button:hover {
  transform: scale(1.2);
  transition: 0.3s ease-out;
}

/* Don't forget to hide the original file input! */
.file-select>input[type="file"] {
  display: none;
}

ul{
  list-style:inside;
}
.errors{
  display: flex;
  gap: 25px;
}

.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>

