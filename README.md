# Blood Donation App

## Overview
The Blood Donation App is a web application designed to facilitate the process of blood donation. It allows users to register as donors, find patients in need of blood, and manage donations efficiently.

## Features
- **Donor Registration**: Users can register as blood donors by providing their personal details and blood type.
- **Patient Registration**: Users can register patients who need blood donations, specifying the required blood type and urgency.
- **Donation Management**: Admins can manage donations, matching donors with patients based on blood type compatibility.
- **WhatsApp Verification**: Verification codes are sent to users via WhatsApp to ensure authenticity.

## Technologies Used
- **Frontend**: HTML, CSS, JavaScript
- **Backend**: Go (Golang), Gin framework
- **Database**: SQLite (or any other preferred database)
- **Other**: Python for WhatsApp integration

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/abdulkarim1422/BloodsApp.git
    ```
2. Navigate to the project directory:
    ```sh
    cd BloodsApp
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```
4. Run the application:
    ```sh
    go run main.go
    ```

## Usage
1. Open your web browser and navigate to `http://localhost:8080`.
2. Use the navigation menu to access different features of the application.
3. Register as a donor or patient, and manage donations through the admin dashboard.

## Contributing
Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License
This project is licensed under the MIT License.


# تطبيق التبرع بالدم

## نظرة عامة
تطبيق التبرع بالدم هو تطبيق ويب مصمم لتسهيل عملية التبرع بالدم. يتيح للمستخدمين التسجيل كمتبرعين، والعثور على المرضى المحتاجين للدم، وإدارة التبرعات بكفاءة.

## الميزات
- **تسجيل المتبرعين**: يمكن للمستخدمين التسجيل كمتبرعين بالدم عن طريق تقديم تفاصيلهم الشخصية وفصيلة الدم.
- **تسجيل المرضى**: يمكن للمستخدمين تسجيل المرضى الذين يحتاجون إلى تبرعات بالدم، مع تحديد فصيلة الدم المطلوبة ودرجة الاستعجال.
- **إدارة التبرعات**: يمكن للمسؤولين إدارة التبرعات، ومطابقة المتبرعين مع المرضى بناءً على توافق فصيلة الدم.
- **التحقق عبر الواتساب**: يتم إرسال رموز التحقق إلى المستخدمين عبر الواتساب لضمان الأصالة.

## التقنيات المستخدمة
- **الواجهة الأمامية**: HTML, CSS, JavaScript
- **الواجهة الخلفية**: Go (Golang), إطار عمل Gin
- **قاعدة البيانات**: SQLite (أو أي قاعدة بيانات مفضلة أخرى)
- **أخرى**: Python للتكامل مع الواتساب

## التثبيت
1. استنساخ المستودع:
    ```sh
    git clone https://github.com/abdulkarim1422/BloodsApp.git
    ```
2. الانتقال إلى دليل المشروع:
    ```sh
    cd BloodsApp
    ```
3. تثبيت التبعيات:
    ```sh
    go mod tidy
    ```
4. تشغيل التطبيق:
    ```sh
    go run main.go
    ```

## الاستخدام
1. افتح متصفح الويب الخاص بك وانتقل إلى `http://localhost:8080`.
2. استخدم قائمة التنقل للوصول إلى الميزات المختلفة للتطبيق.
3. سجل كمتبرع أو مريض، وقم بإدارة التبرعات من خلال لوحة تحكم المسؤول.

## المساهمة
المساهمات مرحب بها! يرجى استنساخ المستودع وإنشاء طلب سحب مع التغييرات الخاصة بك.

## الترخيص
هذا المشروع مرخص بموجب ترخيص MIT.