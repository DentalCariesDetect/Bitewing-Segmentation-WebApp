
from ultralytics import YOLO
from PIL import Image

import cv2
import numpy as np
import os
from datetime import datetime

current_directory = os.path.dirname(__file__)

def save_image(file):
    uploaded_file = file.file.read()
    
    output_folder = os.path.join(current_directory, "img")
    if not os.path.exists(output_folder):
        os.makedirs(output_folder)
    
    current_time = datetime.now()
    time_format = '%Y%m%d-%H%M%S'
    formatted_time = current_time.strftime(time_format)
    output_filename = f'{formatted_time}-ori.jpg'
    print(output_filename)
    output_path = os.path.join(output_folder,output_filename)

    with open(output_path, 'wb') as output_file:
        output_file.write(uploaded_file)
    
    return output_path

def cropping(img_path):
    img_path_list = []
    img_path_crop_list = ""
    
    full_model_path = os.path.join(current_directory, "best.pt")
    model = YOLO(full_model_path)

    result = model(img_path)
    print(result)
    boxes = result[0].plot()
    Image.fromarray(boxes).show()

    file_name, file_extension = os.path.splitext(os.path.basename(img_path))
    img_name = file_name[:-4]
    count = 0
    #classNames = ['U1', 'U2', 'U3', 'U4', 'U5', 'U6', 'L1', 'L2', 'L3', 'L4', 'L5', 'L6', 'Unknown']
    classNames = ['13','14','15','16','17','18','43','44','45','46','47','48','object']
    color = [[0,0,255], [0,255,0], [255,0,0], [0,255,255], [150,255,150], [255,0,255], [24,67,85], [25,55,123], [14,128,128], [128,0,0], [128,128,0], [0,128,0], [0,128,128]]
    bbox_xyxys = (result[0].boxes.xyxy).tolist()
    labels = result[0].boxes.cls.tolist()
    frame = cv2.imread(img_path)
    
    
    for (bbox_xyxy, cls) in zip(bbox_xyxys, labels):
        bbox = np.array(bbox_xyxy)
        x1, y1, x2, y2 = bbox[0], bbox[1], bbox[2], bbox[3]
        x1, y1, x2, y2 = int(x1), int(y1), int(x2), int(y2)

        cropped = frame[y1:y2, x1:x2]
        classname = classNames[int(cls)]

        cv2.imwrite(f'{current_directory}/img' + f'/{img_name}-{classname}' + '.jpg', cropped)
        img_path_list.append(f'{current_directory}/img' + f'/{img_name}-{classname}' + '.jpg')
    for (bbox_xyxy, cls) in zip(bbox_xyxys, labels):
        bbox = np.array(bbox_xyxy)
        x1, y1, x2, y2 = bbox[0], bbox[1], bbox[2], bbox[3]
        x1, y1, x2, y2 = int(x1), int(y1), int(x2), int(y2)

        classname = classNames[int(cls)]
        label = f'{classname}'

        t_size = cv2.getTextSize(label, 0, fontScale=1, thickness=2)[0]
        c2 = x1 + t_size[0] + 30, y1 + t_size[1] + 50
        cv2.rectangle(frame, (x1, y1), c2,color[int(cls)] , -1, cv2.LINE_AA)
        cv2.putText(frame, label, (x1 + 10, y1 + 60), 0, 1.5, [255,255,255] ,thickness=3
                    , lineType=cv2.LINE_AA)
        cv2.rectangle(frame, (x1, y1), (x2, y2), color[int(cls)], 5)
    Image.fromarray(frame).save(f'{current_directory}/img' + f'/{img_name}-overview' + '.jpg')
    img_path = f'{current_directory}/img' + f'/{img_name}-overview' + '.jpg'
    
    return img_path_list, img_path
    
