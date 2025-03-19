INSERT INTO homestay_business (title, user_id, info, boss_info, license_fron, license_back, row_state, star, tags, cover, header_img, version)
VALUES('阳光假日民宿', 7, '温馨舒适，靠近海滩', '李老板，热情好客', 'license1.jpg', 'license2.jpg', 1, 4.8, '海景', 'cover1.jpg', 'header1.jpg', 1);


INSERT INTO homestay (title, sub_title, banner, info, people_num, homestay_business_id, user_id, row_state, row_type, food_info, food_price, homestay_price, market_homestay_price, version)
VALUES
    ('海景房101', '带阳台海景房', 'banner1.jpg', '适合情侣或小家庭', 2, 6, 7, 1, 0, '含早餐', 5000, 20000, 25000, 1),
    ('山景木屋', '宁静木屋体验', 'banner2.jpg', '森林中的小屋，享受安静', 4, 6, 7, 1, 1, '含三餐', 10000, 30000, 35000, 1),
    ('湖畔小筑', '湖边独立木屋', 'banner3.jpg', '宁静湖畔，远离喧嚣', 3, 6, 7, 1, 0, '含早餐和晚餐', 8000, 22000, 28000, 1),
    ('星空露营地', '户外帐篷体验', 'banner4.jpg', '夜晚可观赏满天星斗', 2, 6, 7, 1, 1, '自助烧烤', 7000, 18000, 23000, 1),
    ('田园小屋', '乡村风情，舒适安逸', 'banner5.jpg', '适合家庭出游，享受田园生活', 5, 6, 7, 1, 0, '地道农家饭', 6000, 25000, 30000, 1),
    ('山顶别墅', '360°观景房', 'banner6.jpg', '超大露台，可俯瞰整座山脉', 6, 6, 7, 1, 1, '自助厨房', 0, 50000, 60000, 1),
    ('竹林庭院', '隐世幽居', 'banner7.jpg', '围绕竹林，环境清幽', 3, 6, 7, 1, 0, '素食套餐', 5000, 20000, 25000, 1),
    ('温泉度假屋', '天然温泉泡汤', 'banner8.jpg', '独立温泉池，放松身心', 4, 6, 7, 1, 1, '日式料理', 12000, 35000, 40000, 1),
    ('森林木屋', '童话般的小屋', 'banner9.jpg', '森林环绕，鸟语花香', 2, 6, 7, 1, 0, '欧式早餐', 4000, 15000, 20000, 1),
    ('悬崖海景房', '悬崖边的壮丽景色', 'banner10.jpg', '靠近悬崖，极致海景体验', 2, 6, 7, 1, 1, '海鲜大餐', 15000, 45000, 50000, 1);

INSERT INTO homestay_activity (row_type, data_id, row_status, version)
VALUES
    ('preferredHomestay', 35, 1, 1),
    ('preferredHomestay', 36, 1, 1),
    ('preferredHomestay', 37, 1, 1),
    ('preferredHomestay', 38, 1, 1);

