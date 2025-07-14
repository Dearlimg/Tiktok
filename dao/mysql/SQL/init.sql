set names utf8mb4;
set foreign_key_checks =0;


-- 1. 用户表（无外键依赖，基础表）
CREATE TABLE `user` (
                        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id，自增主键',
                        `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
                        `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
                        `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录创建时间',
                        `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
                        PRIMARY KEY (`id`) USING BTREE,
                        UNIQUE INDEX `name` (`name`) USING BTREE COMMENT '用户名保证唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- 2. 视频表（依赖 user 表的 id 作为外键）
CREATE TABLE `video` (
                         `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '视频id',
                         `author_id` int(11) NOT NULL COMMENT '视频作者id',
                         `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频标题',
                         `play_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频播放地址',
                         `cover_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频上传时封面地址',
                         `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '视频上传时间',
                         `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '视频更新时间',
                         PRIMARY KEY (`id`) USING BTREE,
                         INDEX `fk_video_user` (`author_id`) USING BTREE,
                         CONSTRAINT `fk_video_user` FOREIGN KEY (`author_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- 3. 评论表（依赖 user 表和 video 表的 id 作为外键）
CREATE TABLE `comment` (
                           `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '评论记录 id',
                           `user_id` int(11) NOT NULL COMMENT '发布评论的用户 id',
                           `video_id` int(11) NOT NULL COMMENT '评论视频 id',
                           `content` varchar(255) NOT NULL COMMENT '评论的内容',
                           `action_type` int(11) NOT NULL COMMENT '1 表示已发表评论，2 表示删除评论',
                           `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '评论发布时间',
                           `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '评论更新时间',
                           PRIMARY KEY (`id`) USING BTREE,
                           INDEX `fk_comment_user` (`user_id`) USING BTREE,
                           INDEX `fk_comment_video` (`video_id`) USING BTREE,
                           CONSTRAINT `fk_comment_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                           CONSTRAINT `fk_comment_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- 4. 点赞表（依赖 user 表和 video 表的 id 作为外键）
CREATE TABLE `like` (
                        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '点赞记录 id',
                        `user_id` int(11) NOT NULL COMMENT '点赞用户 id',
                        `video_id` int(11) NOT NULL COMMENT '点赞视频 id',
                        `liked` int(11) NOT NULL COMMENT '默认 1 表示已点赞，0 表示未点赞',
                        `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '点赞记录创建时间',
                        `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '点赞记录更新时间',
                        PRIMARY KEY (`id`) USING BTREE,
                        INDEX `fk_like_user` (`user_id`) USING BTREE,
                        INDEX `fk_like_video` (`video_id`) USING BTREE,
                        CONSTRAINT `fk_like_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                        CONSTRAINT `fk_like_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- 5. 消息表（依赖 user 表的 id 作为外键）
CREATE TABLE `message` (
                           `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'message 记录 id',
                           `user_id` int(11) NOT NULL COMMENT '发送 message 的 user_id',
                           `receiver_id` int(11) NOT NULL COMMENT '接收 message 的 user_id',
                           `msg_content` varchar(255) NOT NULL COMMENT '消息内容',
                           `msg_type` int(11) NOT NULL COMMENT '消息类型，1 表示发送，2 表示撤回',
                           `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '消息创建时间',
                           `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '消息更新时间',
                           PRIMARY KEY (`id`) USING BTREE,
                           INDEX `fk_message_user_1` (`user_id`) USING BTREE,
                           INDEX `fk_message_user_2` (`receiver_id`) USING BTREE,
                           CONSTRAINT `fk_message_user_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                           CONSTRAINT `fk_message_user_2` FOREIGN KEY (`receiver_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- 6. 关系表（依赖 user 表的 id 作为外键）
CREATE TABLE `relation` (
                            `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'follow关系id',
                            `user_id` int(11) NOT NULL COMMENT '用户id',
                            `following_id` int(11) NOT NULL COMMENT 'user_id关注的用户id',
                            `followed` int(11) NOT NULL COMMENT '默认0表示未关注，1表示已关注',
                            `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录创建时间',
                            `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
                            PRIMARY KEY (`id`) USING BTREE,
                            INDEX `fk_relation_user_1` (`user_id`) USING BTREE,
                            INDEX `fk_relation_user_2` (`following_id`) USING BTREE,
                            CONSTRAINT `fk_relation_user_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                            CONSTRAINT `fk_relation_user_2` FOREIGN KEY (`following_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

set foreign_key_checks = 1;