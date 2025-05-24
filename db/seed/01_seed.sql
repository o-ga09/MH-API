-- モンスター
INSERT INTO monster (monster_id, name, element, another_name, created_at, updated_at) VALUES
 (1, "リオレイア","火", "雌火竜", now(), now()),
 (2, "リオレイア亜種","火", "桜火竜", now(), now()),
 (3, "リオレイア希少種","火", "金火竜", now(), now()),
 (4, "紫毒姫リオレイア", "火", "雌火竜", now(), now()),
 (5, "ヌシ・リオレイア","火", "雌火竜（二つ名）", now(), now()),
 (6, "リオレウス", "火","火竜", now(), now()),
 (7, "リオレウス亜種","火", "蒼火竜", now(), now()),
 (8, "リオレウス希少種","火", "銀火竜", now(), now()),
 (9, "黒炎王リオレウス","火", "火竜（二つ名）", now(), now()),
 (10, "ヌシ・リオレウス","火", "火竜（ヌシ）", now(), now());
 -- 種族
INSERT INTO tribe (tribe_id, name_ja, name_en, monster_id, created_at, updated_at) VALUES
 (1, "飛竜種", "Flying Wyvern", 1, now(), now()),
 (2, "飛竜種", "Flying Wyvern", 2, now(), now()),
 (3, "飛竜種", "Flying Wyvern", 3, now(), now()),
 (4, "飛竜種", "Flying Wyvern", 4, now(), now()),
 (5, "飛竜種", "Flying Wyvern", 5, now(), now()),
 (6, "飛竜種", "Flying Wyvern", 6, now(), now()),
 (7, "飛竜種", "Flying Wyvern", 7, now(), now()),
 (8, "飛竜種", "Flying Wyvern", 8, now(), now()),
 (9, "飛竜種", "Flying Wyvern", 9, now(), now()),
 (10, "飛竜種", "Flying Wyvern", 10, now(), now());
 -- 作品
 INSERT INTO product (product_id, name, monster_id, created_at, updated_at) VALUES 
 (1, "MH", 1, now(), now()),
 (2, "MHG", 1, now(), now()),
 (3, "MHP", 1, now(), now()),
 (4, "MH2", 1, now(), now()),
 (5, "P2nd", 1, now(), now()),
 (6, "P2G", 1, now(), now()),
 (7, "P3rd", 1, now(), now()),
 (8, "MH3G", 1, now(), now()),
 (9, "MH3", 1, now(), now()),
 (10, "MH4", 1, now(), now()),
 (11, "MH4G", 1, now(), now()),
 (12, "MHX", 1, now(), now()),
 (13, "MHXX", 1, now(), now()),
 (14, "MHW", 1, now(), now()),
 (15, "MHWI", 1, now(), now()),
 (16, "MHR", 1, now(), now()),
 (17, "MHRS", 1, now(), now()),
 (2, "MHG", 2, now(), now()),
 (3, "MHP", 2, now(), now()),
 (4, "MH2", 2, now(), now()),
 (5, "P2nd", 2, now(), now()),
 (6, "P2G", 2, now(), now()),
 (8, "MH3G", 2, now(), now()),
 (10, "MH4", 2, now(), now()),
 (11, "MH4G", 2, now(), now()),
 (14, "MHW", 2, now(), now()),
 (15, "MHWI", 2, now(), now()),
 (2, "MHG", 3, now(), now()),
 (3, "MHP", 3, now(), now()),
 (4, "MH2", 3, now(), now()),
 (5, "P2nd", 3, now(), now()),
 (6, "P2G", 3, now(), now()),
 (7, "P3rd", 3, now(), now()),
 (8, "MH3G", 3, now(), now()),
 (10, "MH4", 3, now(), now()),
 (11, "MH4G", 3, now(), now()),
 (12, "MHX", 3, now(), now()),
 (13, "MHXX", 3, now(), now()),
 (15, "MHWI", 3, now(), now()),
 (17, "MHRS", 3, now(), now()),
 (12, "MHX", 4, now(), now()),
 (13, "MHXX", 4, now(), now()),
 (16, "MHR", 5, now(), now()),
 (17, "MHRS", 5, now(), now()),
 (1, "MH", 6, now(), now()),
 (2, "MHG", 6, now(), now()),
 (3, "MHP", 6, now(), now()),
 (4, "MH2", 6, now(), now()),
 (5, "P2nd", 6, now(), now()),
 (6, "P2G", 6, now(), now()),
 (7, "P3rd", 6, now(), now()),
 (8, "MH3G", 6, now(), now()),
 (9, "MH3", 6, now(), now()),
 (10, "MH4", 6, now(), now()),
 (11, "MH4G", 6, now(), now()),
 (12, "MHX", 6, now(), now()),
 (13, "MHXX", 6, now(), now()),
 (14, "MHW", 6, now(), now()),
 (15, "MHWI", 6, now(), now()),
 (16, "MHR", 6, now(), now()),
 (17, "MHRS", 6, now(), now()),
 (2, "MHG", 7, now(), now()),
 (3, "MHP", 7, now(), now()),
 (4, "MH2", 7, now(), now()),
 (5, "P2nd", 7, now(), now()),
 (6, "P2G", 7, now(), now()),
 (8, "MH3G", 7, now(), now()),
 (10, "MH4", 7, now(), now()),
 (11, "MH4G", 7, now(), now()),
 (14, "MHW", 7, now(), now()),
 (15, "MHWI", 7, now(), now()),
 (2, "MHG", 8, now(), now()),
 (3, "MHP", 8, now(), now()),
 (4, "MH2", 8, now(), now()),
 (5, "P2nd", 8, now(), now()),
 (6, "P2G", 8, now(), now()),
 (7, "P3rd", 8, now(), now()),
 (8, "MH3G", 8, now(), now()),
 (10, "MH4", 8, now(), now()),
 (11, "MH4G", 8, now(), now()),
 (12, "MHX", 8, now(), now()),
 (13, "MHXX", 8, now(), now()),
 (15, "MHWI", 8, now(), now()),
 (17, "MHRS", 8, now(), now()),
 (12, "MHX", 9, now(), now()),
 (13, "MHXX", 9, now(), now()),
 (16, "MHR", 10, now(), now()),
 (17, "MHRS", 10, now(), now());
-- ランキング
INSERT INTO ranking (ranking, vote_year, monster_id, created_at, updated_at) VALUES 
 (78, "2024", 1, now(), now()),
 (119, "2024", 2, now(), now()),
 (113, "2024", 3, now(), now()),
 (137, "2024", 4, now(), now()),
 (177, "2024", 5, now(), now()),
 (21, "2024", 6, now(), now()),
 (112, "2024", 7, now(), now()),
 (74, "2024", 8, now(), now()),
 (66, "2024", 9, now(), now()),
 (149, "2024", 10, now(), now());
-- BGM
 INSERT INTO music (monster_id, music_id, name, url, created_at, updated_at) VALUES 
 (1,1, "太古の律動/リオレイア", "jLgjOfT_elA", now(), now()),
 (2,2, "太古の律動/リオレイア", "jLgjOfT_elA", now(), now()),
 (3,3, "塔に現る幻/キリン", "u9VKblxtzyQ", now(), now()),
 (4,4, "決意を胸に灯して", "27tXmZCFtzU", now(), now()),
 (5,5, "采邑追われし赤き咆哮", "D5Qp6zUa828", now(), now()),
 (6,6, "咆哮/リオレウス", "R7OgSwgUQSQ", now(), now()),
 (7,7, "咆哮/リオレウス", "R7OgSwgUQSQ", now(), now()),
 (8,8, "塔に現る幻/キリン", "u9VKblxtzyQ", now(), now()),
 (9,9, "決意を胸に灯して", "27tXmZCFtzU", now(), now()),
 (10,10, "采邑追われし赤き咆哮", "D5Qp6zUa828", now(), now());

-- モンスターの武器
INSERT INTO `weapon` 
(`created_at`, `updated_at`, `deleted_at`, `weapon_id`, `name`, `image_url`, 
`rarerity`, `attack`, `element_attack`, `shapness`, `critical`, `description`) 
VALUES
-- 大剣
(NOW(), NOW(), NULL, 'GS001', '鉄刀【朧】', 'https://example.com/images/weapons/gs001.png', 
'10', '1104', '無', 'ABCDEF', '0%', 'モンスターの素材を一切使わない、鉄鉱石のみで作られた基本的な大剣'),

(NOW(), NOW(), NULL, 'GS002', '真・黒龍剣【煉黒】', 'https://example.com/images/weapons/gs002.png', 
'12', '1536', '龍330', 'CDEFGH', '0%', '黒き龍の魂を宿した禍々しき大剣。振るえば大地すら割れる'),

(NOW(), NOW(), NULL, 'GS003', '金色の斬罪【煌然】', 'https://example.com/images/weapons/gs003.png',
'11', '1392', '無', 'BCDEFG', '25%', '金獅子の力を宿した威風堂々とした大剣。圧倒的な攻撃力を誇る'),

-- 太刀
(NOW(), NOW(), NULL, 'LS001', '蒼炎刀【永世】', 'https://example.com/images/weapons/ls001.png',
'11', '990', '火300', 'DEFGHI', '5%', '青き炎を纏った美しい曲線を描く太刀。切れ味は一級品'),

(NOW(), NOW(), NULL, 'LS002', '雷峰【極光】', 'https://example.com/images/weapons/ls002.png',
'12', '924', '雷380', 'CDEFGH', '15%', '雷と共に閃く太刀。一閃すれば稲妻の如き切れ味で敵を両断する'),

(NOW(), NOW(), NULL, 'LS003', '封龍剣【冥界】', 'https://example.com/images/weapons/ls003.png',
'12', '957', '龍420', 'EFGHIJ', '10%', '黒龍の魂を封じ込めた禍々しき刀。使い手の命を糧に力を増す'),

-- 片手剣
(NOW(), NOW(), NULL, 'SNS001', '王牙剣【天嵐】', 'https://example.com/images/weapons/sns001.png',
'10', '350', '無', 'BCDEFG', '0%', '古の獣の牙から作られた片手剣。軽量ながら優れた切れ味を持つ'),

(NOW(), NOW(), NULL, 'SNS002', '炎剣リオレウス', 'https://example.com/images/weapons/sns002.png',
'11', '322', '火280', 'CDEFGH', '10%', '火竜の魂を宿した片手剣。炎の力で敵を焼き尽くす'),

(NOW(), NOW(), NULL, 'SNS003', '氷刃【雪華】', 'https://example.com/images/weapons/sns003.png',
'12', '336', '氷350', 'DEFGHI', '5%', '冷気を放つ刃を持つ片手剣。触れるだけで凍りつかせる'),

-- ランス
(NOW(), NOW(), NULL, 'LNC001', '真・皇金の槍【極天】', 'https://example.com/images/weapons/lnc001.png',
'12', '782', '無', 'BCDEFG', '0%', '黄金の輝きを放つランス。貫通力は群を抜く'),

(NOW(), NOW(), NULL, 'LNC002', '雷電槍【極震】', 'https://example.com/images/weapons/lnc002.png',
'11', '713', '雷320', 'CDEFGH', '0%', '雷の力を宿したランス。突きに雷撃を伴う'),

(NOW(), NOW(), NULL, 'LNC003', '氷結槍【凍土】', 'https://example.com/images/weapons/lnc003.png',
'11', '690', '氷360', 'BCDEFG', '0%', '凍てつく力を宿した槍。大地を凍結させる力を持つ'),

-- ハンマー
(NOW(), NOW(), NULL, 'HMR001', '鬼神轟槌【羅刹】', 'https://example.com/images/weapons/hmr001.png',
'12', '1560', '無', 'ABCDE', '0%', '鬼の力を宿した巨大なハンマー。振るうだけで風を切り裂く'),

(NOW(), NOW(), NULL, 'HMR002', '雷槌【轟天】', 'https://example.com/images/weapons/hmr002.png',
'11', '1404', '雷230', 'BCDEF', '5%', '天の雷を封じ込めたハンマー。一撃ごとに雷鳴が響く'),

(NOW(), NOW(), NULL, 'HMR003', '氷塊槌【凍土】', 'https://example.com/images/weapons/hmr003.png',
'10', '1352', '氷280', 'ABCDE', '0%', '永久凍土の力を宿したハンマー。打撃に凍結の力を伴う'),

-- 狩猟笛
(NOW(), NOW(), NULL, 'HH001', '神韻の竜頭琴', 'https://example.com/images/weapons/hh001.png',
'12', '1152', '無', 'BCDEFG', '0%', '古の竜の頭部を模した狩猟笛。奏でる音は竜をも従わせる'),

(NOW(), NOW(), NULL, 'HH002', '火竜琴【紅蓮】', 'https://example.com/images/weapons/hh002.png',
'11', '1008', '火270', 'CDEFGH', '0%', '火竜の力を宿した狩猟笛。炎の旋律を奏でる'),

(NOW(), NOW(), NULL, 'HH003', '氷結琴【白霜】', 'https://example.com/images/weapons/hh003.png',
'11', '1056', '氷240', 'BCDEFG', '5%', '凍てつく力を宿した狩猟笛。冷気の音色が響き渡る'),

-- 双剣
(NOW(), NOW(), NULL, 'DB001', '真・羅刹双刃【煉獄】', 'https://example.com/images/weapons/db001.png',
'12', '350x2', '火350', 'DEFGHI', '10%', '鬼の力を宿した双剣。炎のような赤い刃が特徴'),

(NOW(), NOW(), NULL, 'DB002', '双雷剣【極光】', 'https://example.com/images/weapons/db002.png',
'11', '322x2', '雷280', 'CDEFGH', '15%', '雷の力を宿した双剣。稲妻の如き斬撃を繰り出す'),

(NOW(), NOW(), NULL, 'DB003', '氷刃双剣【雪風】', 'https://example.com/images/weapons/db003.png',
'12', '336x2', '氷320', 'DEFGHI', '5%', '凍てつく刃を持つ双剣。敵を凍らせる冷気を放つ'),

-- 弓
(NOW(), NOW(), NULL, 'BOW001', '皇金の弓【天嵐】', 'https://example.com/images/weapons/bow001.png',
'12', '390', '無', 'N/A', '15%', '黄金に輝く高貴な弓。射撃の精度は群を抜く'),

(NOW(), NOW(), NULL, 'BOW002', '炎弓【紅蓮】', 'https://example.com/images/weapons/bow002.png',
'11', '360', '火300', 'N/A', '10%', '火の力を宿した弓。矢は炎を纏い敵を焼き尽くす'),

(NOW(), NOW(), NULL, 'BOW003', '氷結弓【白絹】', 'https://example.com/images/weapons/bow003.png',
'12', '372', '氷350', 'N/A', '5%', '凍てつく力を宿した弓。冷気の矢が敵を貫く');
