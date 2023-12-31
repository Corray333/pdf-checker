import pytesseract
from PIL import Image
import cv2
import prepocessing
from pdf2image import convert_from_path
import numpy as np
from flask import Flask, request, jsonify
import json

pytesseract.pytesseract.tesseract_cmd = r'C:/Program Files/Tesseract-OCR/tesseract.exe'

def OCR_crop_info(image, info):
    """
    Обработка зеголовка для конкретного документа,
    формирование словаря для оцифрованных данных.

    :param image: обрезанное изображение зголовков
    :param info: словарь, содержащий пустую строку и координаты текстового блока

    :return: словарь, содержащий оцифрованный текст, находившийся в заданном текстовом поле
    """
    for key in info:
        coord = info[key][1]
        current_data = crop_create(image, coord)
        info[key][0] = current_data.replace("\n", " ")
        info[key][0] = info[key][0][:-1]
        info[key] = info[key][0] # убираю координаты текстового блока
    return info

def crop_create(image, coord):
    """
    Обрезает изображение согласно заданным координатам, выполняет распознавание текста с помощью pytesseract.

    :param image: Исходное изображение, которое будет обрезано и распознано.
    :param coord: Координаты обрезки в формате (x1, y1, x2, y2),
    где (x1, y1) - верхний левый угол, (x2, y2) - нижний правый угол.

    :return: Распознанный текст на обрезанном изображении.
    """
    x1, y1, x2, y2 = coord[0], coord[1], coord[2], coord[3]
    temp_image = image.crop((x1,y1,x2,y2))
    temp_image.save("out.jpg")
    value = pytesseract.image_to_string(temp_image, lang="rus")
    return value


# Конвертация изображения из PIL в CV2
def pil_to_cv2(pil_image):
    """
    Конвертирует изображение из формата Pillow (PIL) в формат OpenCV (cv2).

    :param pil_image: Исходное изображение в формате Pillow (PIL), которое требуется преобразовать.
    :type pil_image: PIL.Image.Image

    :return: Изображение в формате OpenCV (cv2).
    :rtype: numpy.ndarray

    Примечание:
    Функция сохраняет временное изображение "temp.jpg" с использованием Pillow,
    а затем загружает его с помощью OpenCV (cv2) и возвращает результат.
    """
    pil_image.save("temp.jpg")
    cv2_image = cv2.imread("temp.jpg")
    return cv2_image


# Конвертация изображения из CV2 в PIL
def cv2_to_pil(cv2_image):
    """"
    Конвертирует изображение из формата OpenCV (cv2) в формат Pillow (PIL).

    :param cv2_image: Исходное изображение в формате OpenCV (cv2), которое требуется преобразовать.
    :type cv2_image: numpy.ndarray

    :return: Изображение в формате Pillow (PIL).
    :rtype: PIL.Image.Image

    Примечание:
    Функция сохраняет временное изображение "temp.jpg" с использованием OpenCV (cv2), а затем открывает его с помощью Pillow (PIL) и возвращает результат.
    """
    cv2.imwrite("temp.jpg",cv2_image)
    pil_image = Image.open("temp.jpg")
    return pil_image


def preparation(image):
    """
    Выполняет предварительную обработку изображения, включая преобразование формата из Pillow (PIL) в OpenCV (cv2),
    бинаризацию, удаление блоков с знаками и возврат результата в формате Pillow (PIL).

    :param image: Исходное изображение в формате Pillow (PIL), которое требуется предварительно обработать.
    :type image: PIL.Image.Image

    :return: Результат предварительной обработки в формате Pillow (PIL).
    :rtype: PIL.Image.Image

    Примечание:
    - Функция начинает с преобразования изображения из формата Pillow (PIL) в формат OpenCV (cv2).
    - Затем выполняется бинаризация изображения с использованием функции prepocessing.binarization.
    - После этого производится удаление блоков с знаками с использованием функции prepocessing.delete_sign_block.
    - Итоговое изображение в формате OpenCV (cv2) преобразуется обратно в формат Pillow (PIL).
    """
    temp_image = pil_to_cv2(image)
    prepared_image = prepocessing.binarization(temp_image)
    result_image = prepocessing.delete_sign_block(prepared_image)
    result = cv2_to_pil(result_image)
    return result


def OCR_M11(input_path):
    info_M11 = {
        'order': ['', [2380, 420, 2790, 537]],
        'organisation': ['', [845, 750, 3000, 840]],
        'structpod': ['', [870, 850, 2980, 1120]],
        'datacreate': ['', [510, 1530, 800, 1770]],
        'optypecode': ['', [820, 1530, 1080, 1770]],
        'sendertructpod': ['', [1100, 1530, 1550, 1770]],
        'restructpod': ['', [1950, 1530, 2370, 1770]],
        'throughwho': ['', [620, 1850, 2720, 2050]],
        'zatreber': ['', [650, 2100, 2060, 2300]],
        'razr': ['', [2430, 2100, 3830, 2300]]
    }
    """
    словарь, содержащий пустую строку и координаты текстого блока, соответсвующее конкретному полю документа
    """
    pages = convert_from_path(input_path, 500, poppler_path="C:/Program Files/poppler-23.08.0/Library/bin") # конвертируем pdf постранично в jpg

    cropped_image = pages[0].crop((0, 0, 4134, 2350)) # готовим временное изображение для заголовков

    info_heading = OCR_crop_info(cropped_image, info_M11)
    # creating json
    json_file_path = "data.json"

    with open(json_file_path, 'w', encoding="utf-8") as json_file:
        json.dump(info_heading, json_file, ensure_ascii=False)
    print(info_heading)


app = Flask(__name__)

# Определяем декоратор для обработки POST-запросов по пути /post
@app.route('/process', methods=['POST'])
def post_handler():
    # Получаем данные из запроса в формате JSON
    data = request.get_json()
    # Проверяем, что данные не пустые   
    if data:
        # Возвращаем ответ в виде JSON с теми же данными и статусом 200 (успешно)
        return OCR_M11(data["fileName"]), 200
    else:
        # Возвращаем ответ с сообщением об ошибке и статусом 400 (неверный запрос)
        return jsonify({'message': 'No data provided'}), 400

if __name__ == '__main__':
    app.run(host='127.0.0.1', port=3031)


