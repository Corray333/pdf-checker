<template>
  <main>
    <div class="header">
      <span class="logo-container">
        <img src="./assets/logo.png" alt="" class="logo">
        <div>
          <p>OsaSoft</p>
          <p>PDF Checker</p>
        </div>
      </span>
    </div>
    <div class="content">
      <label class="file-select" v-if="file == null">
        <h1>Загрузите файл</h1>
        <!-- We can't use a normal button element here, as it would become the target of the label. -->
        <div class="select-button">
          <!-- Display the filename if a file has been selected. -->
          <span>+</span>
        </div>
        <!-- Now, the file input that we hide. -->
        <input id="fileInput" type="file" @change="uploadFile" />
      </label>
      <object v-else data="https://www.w3docs.com/uploads/media/default/0001/01/540cb75550adf33f281f29132dddd14fded85bfc.pdf" type="application/pdf"
        width="100%" height="500px">
        <p>Unable to display PDF file. <a
            href="/uploads/media/default/0001/01/540cb75550adf33f281f29132dddd14fded85bfc.pdf">Download</a> instead.</p>
      </object>
    </div>
  </main>
</template>

<script setup>
import { ref } from "vue"
// import pdf from 'pdfvuer'

const filePath = ref(null)

function uploadFile() {
  // Получаем элемент input
  let input = document.getElementById("fileInput")
  // Проверяем, что файл выбран
  if (input.files.length > 0) {
    // Создаем объект FormData
    let formData = new FormData()
    // Добавляем файл в объект FormData
    formData.append("myFile", input.files[0])
    filePath.value = input.files[0].filePath
    console.log(filePath.value)
    console.log(input.files[0])
    // Создаем объект XMLHttpRequest
    let xhr = new XMLHttpRequest()
    // Открываем соединение с сервером
    xhr.open("POST", "/test")
    // Отправляем запрос с данными формы
    xhr.send(formData)
    // Обрабатываем ответ сервера
    xhr.onload = function () {
      if (xhr.status == 200) {
        // Выводим сообщение об успешной загрузке
        alert("Файл успешно загружен на сервер")
      } else {
        // Выводим сообщение об ошибке
        alert("Произошла ошибка при загрузке файла: " + xhr.statusText)
      }
    }
  } else {
    // Выводим сообщение, если файл не выбран
    alert("Пожалуйста, выберите файл для загрузки")
  }
}


</script>

<style scoped>
main {
  padding: 0 250px;
  height: 100vh;
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
</style>

