create database home_nest_travel;
use home_nest_travel;

DROP TABLE IF EXISTS `homestay`;
CREATE TABLE `homestay` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `sub_title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '副标题',
  `banner` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '轮播图，第一张封面',
  `info` varchar(4069) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '介绍',
  `people_num` tinyint(1) NOT NULL DEFAULT '0' COMMENT '容纳人的数量',
  `homestay_business_id` bigint NOT NULL DEFAULT '0' COMMENT '民宿店铺id',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '房东id，冗余字段',
  `row_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:下架 1:上架',
  `row_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '售卖类型0：按房间出售 1:按人次出售',
  `food_info` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '餐食标准',
  `food_price` bigint NOT NULL DEFAULT '0' COMMENT '餐食价格（分）',
  `homestay_price` bigint NOT NULL DEFAULT '0' COMMENT '民宿价格（分）',
  `market_homestay_price` bigint NOT NULL DEFAULT '0' COMMENT '民宿市场价格（分）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='每一间民宿';

DROP TABLE IF EXISTS `homestay_activity`;
CREATE TABLE `homestay_activity` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0',
  `row_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动类型',
  `data_id` bigint NOT NULL DEFAULT '0' COMMENT '业务表id（id跟随活动类型走）',
  `row_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:下架 1:上架',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  PRIMARY KEY (`id`),
  KEY `idx_rowType` (`row_type`,`row_status`,`del_state`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='每一间民宿';

DROP TABLE IF EXISTS `homestay_business`;
CREATE TABLE `homestay_business` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0',
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺名称',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '关联的用户id',
  `info` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺介绍',
  `boss_info` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '房东介绍',
  `license_fron` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '营业执照正面',
  `license_back` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '营业执照背面',
  `row_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:禁止营业 1:正常营业',
  `star` double(2,1) NOT NULL DEFAULT '0.0' COMMENT '店铺整体评价，冗余',
  `tags` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '每个店家一个标签，自己编辑',
  `cover` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '封面图',
  `header_img` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店招门头图片',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_userId` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='民宿店铺';

-- app/travel/model
-- goctl model mysql datasource --dir . --table homestay --cache true --url "root:password@tcp(127.0.0.1:3306)/home_nest_travel"
-- goctl model mysql datasource --dir . --table homestay_activity --cache true --url "root:password@tcp(127.0.0.1:3306)/home_nest_travel"
-- goctl model mysql datasource --dir . --table homestay_business --cache true --url "root:password@tcp(127.0.0.1:3306)/home_nest_travel"


INSERT INTO `homestay_business` (
    `user_id`,
    `title`,
    `info`,
    `boss_info`,
    `license_fron`,
    `license_back`,
    `row_state`,
    `star`,
    `tags`,
    `cover`,
    `header_img`,
    `version`
) VALUES
(2, '山水雅居', '坐落于山水之间的宁静民宿，远离城市喧嚣', '热爱自然的80后创业者，喜欢结交各地朋友', '/license/front_1.jpg', '/license/back_1.jpg', 1, 4.5, '自然风光,安静', '/cover/landscape_1.jpg', '/header/mountain_1.jpg', 1),
(3, '城市花园客栈', '位于市中心的花园式民宿，交通便利', '资深旅游爱好者，从事酒店行业10年', '/license/front_2.jpg', '/license/back_2.jpg', 1, 4.2, '市中心,花园', '/cover/city_1.jpg', '/header/garden_1.jpg', 1),
(4, '海边小筑', '面朝大海，春暖花开的海景民宿', '海边长大的本地人，熟悉当地风土人情', '/license/front_3.jpg', '/license/back_3.jpg', 1, 4.8, '海景,浪漫', '/cover/sea_1.jpg', '/header/beach_1.jpg', 1),
(5, '古镇风情', '保留传统建筑风格的古镇民宿', '传统文化保护者，擅长讲解当地历史', '/license/front_4.jpg', '/license/back_4.jpg', 1, 4.3, '古镇,传统', '/cover/ancient_1.jpg', '/header/town_1.jpg', 1),
(6, '竹林幽境', '隐藏在竹林中的禅意民宿', '瑜伽教练，追求心灵平静的生活方式', '/license/front_5.jpg', '/license/back_5.jpg', 1, 4.6, '禅意,养生', '/cover/bamboo_1.jpg', '/header/zen_1.jpg', 1),
(7, '星空观景台', '拥有绝佳星空观测条件的山顶民宿', '天文爱好者，配备专业望远镜', '/license/front_6.jpg', '/license/back_6.jpg', 1, 4.7, '星空,山顶', '/cover/stars_1.jpg', '/header/observatory_1.jpg', 1),
(8, '田园牧歌', '体验农家生活的田园民宿', '退休教师，热爱农耕生活', '/license/front_7.jpg', '/license/back_7.jpg', 1, 4.1, '田园,农家', '/cover/farm_1.jpg', '/header/country_1.jpg', 1),
(9, '艺术工作室', '结合艺术创作与住宿的独特空间', '青年艺术家，经常举办艺术沙龙', '/license/front_8.jpg', '/license/back_8.jpg', 1, 4.4, '艺术,创意', '/cover/art_1.jpg', '/header/studio_1.jpg', 1),
(10, '温泉度假屋', '拥有天然温泉的休闲度假民宿', '温泉疗法师，精通养生之道', '/license/front_9.jpg', '/license/back_9.jpg', 1, 4.9, '温泉,度假', '/cover/hotspring_1.jpg', '/header/spa_1.jpg', 1),
(11, '书香门第', '以书籍为主题的文艺民宿', '前图书编辑，藏书丰富', '/license/front_10.jpg', '/license/back_10.jpg', 1, 4.0, '文艺,书籍', '/cover/books_1.jpg', '/header/library_1.jpg', 1),
(12, '极简生活馆', '倡导极简主义生活方式的民宿', '极简生活博主，著有相关书籍', '/license/front_11.jpg', '/license/back_11.jpg', 1, 4.2, '极简,环保', '/cover/minimal_1.jpg', '/header/simple_1.jpg', 1),
(13, '美食探索家', '以美食体验为特色的民宿', '美食作家，擅长各地菜系', '/license/front_12.jpg', '/license/back_12.jpg', 1, 4.7, '美食,烹饪', '/cover/food_1.jpg', '/header/kitchen_1.jpg', 1),
(14, '运动健康屋', '适合运动爱好者的健康民宿', '健身教练，提供健身指导', '/license/front_13.jpg', '/license/back_13.jpg', 0, 4.3, '运动,健康', '/cover/sports_1.jpg', '/header/gym_1.jpg', 1),
(15, '宠物友好之家', '欢迎携带宠物入住的民宿', '动物保护志愿者，养有多只宠物', '/license/front_14.jpg', '/license/back_14.jpg', 1, 4.5, '宠物友好', '/cover/pets_1.jpg', '/header/animals_1.jpg', 1),
(16, '商务便捷居', '适合商务人士的现代化民宿', '前企业高管，了解商务需求', '/license/front_15.jpg', '/license/back_15.jpg', 1, 4.1, '商务,便捷', '/cover/business_1.jpg', '/header/office_1.jpg', 1);

INSERT INTO `homestay` (
    `title`,
    `sub_title`,
    `banner`,
    `info`,
    `people_num`,
    `homestay_business_id`,
    `user_id`,
    `row_state`,
    `row_type`,
    `food_info`,
    `food_price`,
    `homestay_price`,
    `market_homestay_price`
) VALUES
-- 山水雅居 (user_id:2, business_id:1)
('山水观景房', '180度全景山景，带独立阳台', '/banner/mountain1_1.jpg,/banner/mountain1_2.jpg,/banner/mountain1_3.jpg', '房间面朝群山，早晨可观赏云海，配备现代化设施，独立卫浴，免费WiFi', 2, 1, 2, 1, 0, '提供当地特色早餐，包含新鲜水果、手工面包和现磨咖啡', 5000, 29800, 35800),
('竹林雅舍', '隐藏在竹林中的静谧空间', '/banner/bamboo1_1.jpg,/banner/bamboo1_2.jpg', '被竹林环绕的独立小屋，私密性好，适合冥想和放松', 2, 1, 2, 1, 0, '素食早餐，有机食材', 4500, 25800, 30800),

-- 城市花园客栈 (user_id:3, business_id:2)
('都市花园房', '市中心的花园景观房', '/banner/city1_1.jpg,/banner/city1_2.jpg,/banner/city1_3.jpg', '位于市中心但安静舒适，步行可达商业区，配备齐全', 2, 2, 3, 1, 0, '西式早餐，可送餐到房', 6000, 32800, 38800),
('商务大床房', '适合商务人士的便捷住宿', '/banner/business1_1.jpg,/banner/business1_2.jpg', '配备办公桌和高速网络，临近地铁站', 1, 2, 3, 1, 0, '快捷早餐套餐', 3500, 22800, 27800),

-- 海边小筑 (user_id:4, business_id:3)
('海景套房', '无敌海景，私人露台', '/banner/sea1_1.jpg,/banner/sea1_2.jpg,/banner/sea1_3.jpg', '直面大海，听涛入眠，露台可观赏日出日落', 2, 3, 4, 1, 0, '海鲜早餐，当地特色', 7500, 45800, 52800),
('沙滩小屋', '步行30秒到沙滩', '/banner/beach1_1.jpg,/banner/beach1_2.jpg', '传统渔村风格，体验当地生活', 3, 3, 4, 1, 0, '渔家早餐，新鲜捕捞', 5000, 32800, 39800),

-- 古镇风情 (user_id:5, business_id:4)
('明清古宅', '300年历史古建筑', '/banner/ancient1_1.jpg,/banner/ancient1_2.jpg', '保存完好的明清建筑，体验古代文人生活', 2, 4, 5, 1, 0, '传统茶点，古法制作', 4000, 27800, 33800),
('庭院景观房', '传统中式庭院景观', '/banner/courtyard1_1.jpg,/banner/courtyard1_2.jpg', '围绕中央庭院的客房，四季景致不同', 2, 4, 5, 1, 0, '中式早餐，现做点心', 4500, 29800, 35800),

-- 竹林幽境 (user_id:6, business_id:5)
('禅意单间', '适合单人修行的静谧空间', '/banner/zen1_1.jpg,/banner/zen1_2.jpg', '简约设计，提供冥想坐垫和茶具', 1, 5, 6, 1, 0, '养生粥品，药膳调理', 3000, 19800, 24800),
('双人禅房', '双人修行空间', '/banner/zen2_1.jpg,/banner/zen2_2.jpg', '适合双人同修，共享宁静时光', 2, 5, 6, 1, 0, '双人养生套餐', 5500, 25800, 31800),

-- 星空观景台 (user_id:7, business_id:6)
('天文观测房', '配备专业天文望远镜', '/banner/stars1_1.jpg,/banner/stars1_2.jpg', '屋顶可开启，专业级天文望远镜，星空导师指导', 2, 6, 7, 1, 0, '夜间观星茶点', 4000, 37800, 44800),
('星空露营帐篷', '山顶露营体验', '/banner/camping1_1.jpg,/banner/camping1_2.jpg', '豪华露营装备，篝火晚会', 4, 6, 7, 1, 1, '烧烤晚餐+早餐', 8000, 19800, 25800),

-- 田园牧歌 (user_id:8, business_id:7)
('农家小院', '独立农家院落', '/banner/farm1_1.jpg,/banner/farm1_2.jpg', '带小菜园，可体验农耕乐趣', 4, 7, 8, 1, 0, '农家自产食材', 3500, 32800, 39800),
('谷仓改造房', '旧谷仓改造的特色房间', '/banner/barn1_1.jpg,/banner/barn1_2.jpg', '保留原始结构，现代舒适设施', 2, 7, 8, 1, 0, '农家传统早餐', 3000, 22800, 28800),

-- 艺术工作室 (user_id:9, business_id:8)
('画家工作室', '可创作的艺术空间', '/banner/art1_1.jpg,/banner/art1_2.jpg', '提供画架和基础画材，艺术氛围浓厚', 2, 8, 9, 1, 0, '创意摆盘早餐', 5000, 29800, 35800),
('雕塑主题房', '雕塑艺术品陈列', '/banner/sculpture1_1.jpg,/banner/sculpture1_2.jpg', '房间内有多件原创雕塑作品', 2, 8, 9, 1, 0, '艺术主题餐食', 5500, 32800, 38800),

-- 温泉度假屋 (user_id:10, business_id:9)
('私汤庭院房', '独立温泉泡池', '/banner/hotspring1_1.jpg,/banner/hotspring1_2.jpg', '每间房都有私人温泉池，24小时天然温泉', 2, 9, 10, 1, 0, '日式会席料理', 12000, 68800, 78800),
('家庭温泉套房', '适合家庭的温泉套房', '/banner/familyhotspring1_1.jpg,/banner/familyhotspring1_2.jpg', '两间卧室，共享大温泉池', 4, 9, 10, 1, 0, '家庭套餐', 18000, 98800, 108800),

-- 书香门第 (user_id:11, business_id:10)
('图书馆套房', '三面书墙的阅读空间', '/banner/library1_1.jpg,/banner/library1_2.jpg', '藏书千册，舒适阅读区', 2, 10, 11, 1, 0, '英式下午茶', 4500, 27800, 33800),
('作家书房', '仿知名作家书房设计', '/banner/study1_1.jpg,/banner/study1_2.jpg', '复古书桌，创作灵感源泉', 1, 10, 11, 1, 0, '简餐轻食', 3500, 19800, 24800),

-- 极简生活馆 (user_id:12, business_id:11)
('极简单人间', '最少物品的纯净空间', '/banner/minimal1_1.jpg,/banner/minimal1_2.jpg', '仅必需物品，帮助断舍离', 1, 11, 12, 1, 0, '轻断食早餐', 2500, 15800, 19800),
('冥想空间', '全白设计冥想室', '/banner/meditation1_1.jpg,/banner/meditation1_2.jpg', '隔音良好，适合冥想练习', 1, 11, 12, 1, 0, '能量果汁', 2000, 12800, 16800),

-- 美食探索家 (user_id:13, business_id:12)
('厨师主题房', '可参与烹饪的厨房套房', '/banner/kitchen1_1.jpg,/banner/kitchen1_2.jpg', '专业厨房设备，烹饪课程', 2, 12, 13, 1, 0, '烹饪课程+食材', 15000, 42800, 49800),
('美食家套房', '美食书籍和厨具陈列', '/banner/gourmet1_1.jpg,/banner/gourmet1_2.jpg', '收藏各种烹饪书籍和特色厨具', 2, 12, 13, 1, 0, '主厨定制套餐', 20000, 52800, 59800),

-- 运动健康屋 (user_id:14, business_id:13)
('健身主题房', '室内健身设备', '/banner/gym1_1.jpg,/banner/gym1_2.jpg', '配备基础健身器材，运动氛围', 2, 13, 14, 0, 0, '健身营养餐', 6000, 24800, 30800),
('瑜伽冥想室', '专业瑜伽练习空间', '/banner/yoga1_1.jpg,/banner/yoga1_2.jpg', '木地板，瑜伽垫，冥想角落', 1, 13, 14, 1, 0, '瑜伽养生餐', 4500, 19800, 24800),

-- 宠物友好之家 (user_id:15, business_id:14)
('宠物套房', '带宠物设施的专用房间', '/banner/pet1_1.jpg,/banner/pet1_2.jpg', '宠物床、食盆、玩具齐全', 2, 14, 15, 1, 0, '人宠共享早餐', 5000, 27800, 33800),
('庭院宠物房', '带封闭庭院的宠物房', '/banner/petyard1_1.jpg,/banner/petyard1_2.jpg', '安全庭院，宠物可自由活动', 3, 14, 15, 1, 0, '定制宠物餐', 6000, 35800, 41800),

-- 商务便捷居 (user_id:16, business_id:15)
('行政套房', '商务会客功能齐全', '/banner/executive1_1.jpg,/banner/executive1_2.jpg', '会议桌，高速网络，打印服务', 2, 15, 16, 1, 0, '商务简餐', 7000, 42800, 48800),
('便捷单人间', '经济实惠的商务选择', '/banner/economy1_1.jpg,/banner/economy1_2.jpg', '基础设施齐全，性价比高', 1, 15, 16, 1, 0, '快捷早餐', 3000, 16800, 21800),

-- 继续补充更多房间...
('家庭亲子房', '卡通主题儿童喜欢', '/banner/kids1_1.jpg,/banner/kids1_2.jpg', '儿童床，玩具角，安全设计', 4, 17, 3, 1, 0, '儿童营养餐', 8000, 38800, 45800),
('摄影主题房', '专业摄影背景布置', '/banner/photo1_1.jpg,/banner/photo1_2.jpg', '多个摄影场景，专业灯光', 2, 18, 4, 1, 0, '创意造型餐', 5500, 32800, 38800),
('茶道体验房', '传统茶室设计', '/banner/tea1_1.jpg,/banner/tea1_2.jpg', '茶具齐全，可体验茶道', 2, 19, 5, 1, 0, '茶点套餐', 4000, 25800, 31800),
('复古怀旧房', '80年代怀旧风格', '/banner/retro1_1.jpg,/banner/retro1_2.jpg', '老物件陈列， nostalgic氛围', 2, 20, 6, 1, 0, '怀旧小吃', 3500, 22800, 28800),
('现代设计房', '极简工业风格', '/banner/modern1_1.jpg,/banner/modern1_2.jpg', '设计师家具，艺术感强烈', 2, 21, 7, 1, 0, '现代 fusion料理', 6500, 35800, 42800),
('山林观景房', '全景森林景观', '/banner/forest1_1.jpg,/banner/forest1_2.jpg', '落地窗，森林氧吧', 2, 22, 8, 1, 0, '山野食材', 4500, 27800, 33800),
('城市景观房', '高层城市夜景', '/banner/cityview1_1.jpg,/banner/cityview1_2.jpg', '落地窗，城市天际线', 2, 23, 9, 1, 0, '城市特色早餐', 5000, 32800, 38800),
('农家体验房', '土炕传统体验', '/banner/farmstay1_1.jpg,/banner/farmstay1_2.jpg', '传统土炕，农家装饰', 3, 24, 10, 1, 0, '农家饭菜', 4000, 19800, 25800),
('音乐主题房', '音乐器材陈列', '/banner/music1_1.jpg,/banner/music1_2.jpg', '吉他、键盘等乐器', 2, 25, 11, 1, 0, '音乐主题甜点', 4500, 26800, 32800),
('会议功能房', '小型会议设施', '/banner/meeting1_1.jpg,/banner/meeting1_2.jpg', '投影仪，白板，会议桌', 6, 26, 12, 0, 0, '会议茶歇', 10000, 48800, 55800),
('浪漫套房', '玫瑰花瓣布置', '/banner/romantic1_1.jpg,/banner/romantic1_2.jpg', '心形浴缸，浪漫装饰', 2, 27, 13, 1, 0, '情侣套餐', 12000, 58800, 65800),
('学生宿舍房', '经济实惠上下铺', '/banner/dorm1_1.jpg,/banner/dorm1_2.jpg', '4人间，共享卫浴', 4, 28, 14, 1, 1, '学生优惠餐', 2500, 9800, 12800),
('康养套房', '无障碍设计', '/banner/health1_1.jpg,/banner/health1_2.jpg', '适老化设施，医疗呼叫', 2, 29, 15, 1, 0, '药膳调理', 6000, 32800, 38800),
('背包客床位', '青年旅舍风格', '/banner/hostel1_1.jpg,/banner/hostel1_2.jpg', '8人间，储物柜', 1, 30, 16, 1, 1, '简易早餐', 1500, 5800, 8800),
('海景双床房', '适合朋友出行', '/banner/seaview2_1.jpg,/banner/seaview2_2.jpg', '两张单人床，海景阳台', 2, 3, 4, 1, 0, '双人海鲜早餐', 9000, 38800, 45800),
('家庭海景套房', '两间卧室海景房', '/banner/familysea1_1.jpg,/banner/familysea1_2.jpg', '主卧+次卧，客厅海景', 5, 3, 4, 1, 0, '家庭海鲜盛宴', 20000, 88800, 98800),
('商务双床房', '两人商务出行', '/banner/businesstwin1_1.jpg,/banner/businesstwin1_2.jpg', '两张标准床，工作区', 2, 15, 16, 1, 0, '双人商务餐', 12000, 36800, 42800),
('豪华大床房', 'king size大床', '/banner/deluxe1_1.jpg,/banner/deluxe1_2.jpg', '2米大床，豪华卫浴', 2, 9, 10, 1, 0, '豪华早餐', 8000, 48800, 55800),
('经济单人间', '基础住宿需求', '/banner/budget1_1.jpg,/banner/budget1_2.jpg', '简单舒适，性价比高', 1, 2, 3, 1, 0, '基础早餐', 2500, 12800, 16800),
('特色土窑房', '当地特色建筑', '/banner/cave1_1.jpg,/banner/cave1_2.jpg', '冬暖夏凉，独特体验', 2, 4, 5, 1, 0, '窑烤面包', 4000, 22800, 28800);

INSERT INTO `homestay_activity` (
    `row_type`,
    `data_id`,
    `row_status`,
    `version`
) VALUES
-- preferredHomestay 类型（推荐民宿）
('preferredHomestay', 1, 1, 1),   -- 山水观景房
('preferredHomestay', 3, 1, 1),   -- 海景套房
('preferredHomestay', 5, 1, 1),   -- 禅意单间
('preferredHomestay', 7, 1, 1),   -- 天文观测房
('preferredHomestay', 9, 1, 1),   -- 私汤庭院房
('preferredHomestay', 11, 1, 1),  -- 图书馆套房
('preferredHomestay', 13, 1, 1),  -- 极简单人间
('preferredHomestay', 15, 1, 1),  -- 厨师主题房

-- goodBusiness 类型（生意好的民宿）
('goodBusiness', 2, 1, 1),        -- 都市花园房
('goodBusiness', 4, 1, 1),        -- 沙滩小屋
('goodBusiness', 6, 1, 1),        -- 星空露营帐篷
('goodBusiness', 8, 1, 1),        -- 谷仓改造房
('goodBusiness', 10, 1, 1),       -- 家庭温泉套房
('goodBusiness', 12, 1, 1),       -- 作家书房
('goodBusiness', 14, 1, 1),       -- 冥想空间

-- 混合状态示例（包含下架状态）
('preferredHomestay', 16, 0, 1),  -- 健身主题房（下架状态）
('goodBusiness', 17, 0, 1);       -- 瑜伽冥想室（下架状态）