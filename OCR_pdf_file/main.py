import pytesseract
from PIL import Image
import cv2

import prepocessing
from prepocessing import *

image_file = "input.jpg"
img = cv2.imread(image_file)
gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
ret, thresh = cv2.threshold(gray, 127, 255, 0)
# gray = prepocessing.binarization(img)
contours, _ = cv2.findContours(thresh, cv2.RETR_LIST, cv2.CHAIN_APPROX_SIMPLE)

# Отрисовка контуров
cv2.drawContours(img, contours, -1, (0, 255, 0), 2)

# Отображение изображения с контурами
cv2.imshow('Contours', img)
cv2.waitKey(0)
cv2.destroyAllWindows()

# # Распознавание текста
# for contour in contours:
#     x, y, w, h = cv2.boundingRect(contour)
#     text_region = img[y:y+h, x:x+w]
#     text = pytesseract.image_to_string(text_region, lang="rus")
#     print(text)