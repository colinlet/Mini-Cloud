-- phpMyAdmin SQL Dump
-- version phpStudy 2014
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Jan 09, 2019 at 02:13 PM
-- Server version: 5.5.53
-- PHP Version: 5.4.45

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `cloud`
--

-- --------------------------------------------------------

--
-- Table structure for table `mc_address`
--

CREATE TABLE IF NOT EXISTS `mc_address` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mid` int(11) NOT NULL,
  `user_name` varchar(32) NOT NULL,
  `postal_code` varchar(32) NOT NULL,
  `province_name` varchar(32) NOT NULL,
  `city_name` varchar(32) NOT NULL,
  `county_name` varchar(32) NOT NULL,
  `detail_info` varchar(128) NOT NULL,
  `national_code` varchar(32) NOT NULL,
  `tel_number` varchar(32) NOT NULL,
  `is_use` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='地址表' AUTO_INCREMENT=7 ;

--
-- Dumping data for table `mc_address`
--

INSERT INTO `mc_address` (`id`, `mid`, `user_name`, `postal_code`, `province_name`, `city_name`, `county_name`, `detail_info`, `national_code`, `tel_number`, `is_use`) VALUES
(5, 2, '张三', '100010', '广东省', '深圳市', '宝安区', '新安街道建安一路99号', '110101', '0755-33111111', 0),
(6, 2, '凌枫', '00000', '广东省', '深圳市', '南山区', '西丽湖路4065号(西丽湖东侧)', '440305', '0755-26622888', 1);

-- --------------------------------------------------------

--
-- Table structure for table `mc_category`
--

CREATE TABLE IF NOT EXISTS `mc_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父ID',
  `name` varchar(32) NOT NULL COMMENT '分类名称',
  `img` varchar(128) NOT NULL COMMENT '图片',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='商品分类' AUTO_INCREMENT=32 ;

--
-- Dumping data for table `mc_category`
--

INSERT INTO `mc_category` (`id`, `pid`, `name`, `img`, `status`) VALUES
(1, 0, '护肤品', '', 1),
(2, 0, '彩妆', '', 1),
(3, 0, '美发护发', '', 1),
(4, 0, '口腔护理', '', 1),
(5, 0, '身体女性', '', 1),
(6, 0, '男士护肤', '', 1),
(7, 0, '手表', '', 1),
(8, 0, '包包', '', 1),
(9, 0, '药品', '', 0),
(10, 1, '面膜', 'cate_001.png', 1),
(11, 1, '护肤套装', 'cate_002.png', 1),
(12, 1, '卸妆', 'cate_003.png', 1),
(13, 1, '补水', 'cate_004.png', 1),
(14, 2, '时尚彩妆', 'cate_005.png', 1),
(15, 2, '香水', 'cate_006.png', 1),
(16, 2, 'BB霜', 'cate_007.png', 1),
(17, 2, '口红', 'cate_008.png', 1),
(18, 2, '气垫bb', 'cate_009.png', 1),
(19, 2, '睫毛膏', 'cate_010.png', 1),
(20, 3, '洗发水', 'cate_011.png', 1),
(22, 3, '洗护套装', 'cate_012.png', 1),
(23, 4, '牙膏', 'cate_013.png', 1),
(24, 4, '牙刷', 'cate_014.png', 1),
(25, 4, '漱口水', 'cate_015.png', 1),
(26, 4, '牙贴', 'cate_016.png', 1),
(27, 5, '沐浴露', 'cate_017.png', 1),
(28, 5, '洗手液', 'cate_018.png', 1),
(29, 6, '洁面', 'cate_019.png', 1),
(30, 6, '爽肤水', 'cate_020.png', 1),
(31, 6, '控油', 'cate_021.png', 1);

-- --------------------------------------------------------

--
-- Table structure for table `mc_goods`
--

CREATE TABLE IF NOT EXISTS `mc_goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL COMMENT '商品名称',
  `img` varchar(128) NOT NULL COMMENT '封面图片',
  `desc` varchar(32) NOT NULL COMMENT '商品描述',
  `price` int(11) NOT NULL COMMENT '商品价格',
  `origin_price` int(11) NOT NULL COMMENT '商品原价',
  `category_id` int(11) NOT NULL COMMENT '分类ID',
  `showcase` varchar(240) NOT NULL COMMENT '商品详情轮播图',
  `list` text NOT NULL COMMENT '商品详情图片介绍',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `mark` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='商品表' AUTO_INCREMENT=5 ;

--
-- Dumping data for table `mc_goods`
--

INSERT INTO `mc_goods` (`id`, `name`, `img`, `desc`, `price`, `origin_price`, `category_id`, `showcase`, `list`, `status`, `mark`) VALUES
(1, 'JAYJUN韩国樱花面膜', '001/img.png', '提亮肤色 补水 保湿 滋润', 169, 248, 10, '001/sc1.png,001/sc2.png,001/sc3.png', '001/01.jpg,001/02.jpg,001/03.jpg,001/04.jpg,001/05.jpg,001/06.jpg,001/07.jpg,001/08.jpg,001/09.jpg,001/10.jpg,001/11.jpg,001/12.jpg,001/13.jpg,001/14.jpg', 1, 'https://detail.tmall.hk/hk/item.htm?spm=a220m.1000858.1000725.25.39df5dedkLGUE5&id=537118267969&skuId=3207895064628&areaId=440300&sku=21299:27019&standard=1&user_id=2956756848&cat_id=50029231&is_b=1&rn=1b4dbb6a0ab970cb6442f1d44d3312f3'),
(2, '雪花秀护肤礼盒套装', '002/img.png', '滋盈肌本润颜水乳套装', 840, 1287, 11, '002/sc1.png,002/sc2.png,002/sc3.png', '002/01.jpg,002/02.jpg,002/03.jpg,002/04.jpg,002/05.jpg,002/06.jpg,002/07.jpg,002/08.jpg,002/09.jpg', 1, 'https://detail.tmall.com/item.htm?spm=a220m.1000858.1000725.1.f535a9a77DZ8mG&id=527067530404&skuId=3136038178198&areaId=440300&sku=21299:27019&standard=1&user_id=2738112600&cat_id=50031543&is_b=1&rn=1112e0ebb19b4559ad3c7536e24ecdb1'),
(3, '法国贝德玛卸妆水', '003/img.png', '深层清洁 中性/干性肤质', 138, 199, 12, '003/sc1.jpg,003/sc2.jpg,003/sc3.jpg', '003/01.jpg,003/02.jpg,003/03.jpg,003/04.jpg,003/05.jpg,003/06.jpg,003/07.jpg,003/08.jpg', 1, 'https://detail.tmall.com/item.htm?spm=a220m.1000858.1000725.71.402e4a508PiowB&id=559277164250&skuId=3662305330077&areaId=440300&standard=1&user_id=2807304908&cat_id=50103227&is_b=1&rn=152f8c139263c9d5884d8d193ed1c888'),
(4, '兰芝补水保湿夜间修护', '004/img.png', '水润滋养 温和细腻', 149, 208, 13, '004/sc1.jpg,004/sc2.jpg', '004/01.jpg,004/02.jpg,004/03.jpg,004/04.jpg', 1, 'https://detail.tmall.hk/hk/item.htm?spm=a220m.1000858.1000725.116.4bed1e2fqyn5nL&id=552400657980&skuId=3559972795120&areaId=440300&standard=1&user_id=2549841410&cat_id=50026502&is_b=1&rn=25531bd826349d61fc3f25d238646f53');

-- --------------------------------------------------------

--
-- Table structure for table `mc_order`
--

CREATE TABLE IF NOT EXISTS `mc_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单表' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `mc_slideshow`
--

CREATE TABLE IF NOT EXISTS `mc_slideshow` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `img` varchar(128) NOT NULL,
  `goods_id` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='轮播图' AUTO_INCREMENT=5 ;

--
-- Dumping data for table `mc_slideshow`
--

INSERT INTO `mc_slideshow` (`id`, `img`, `goods_id`, `status`) VALUES
(1, 'slideshow_02.png', 1, 1),
(2, 'slideshow_01.png', 2, 1),
(3, 'slideshow_03.png', 3, 1),
(4, 'slideshow_04.png', 4, 1);

-- --------------------------------------------------------

--
-- Table structure for table `mc_user`
--

CREATE TABLE IF NOT EXISTS `mc_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(32) NOT NULL COMMENT 'openid',
  `session` varchar(32) NOT NULL COMMENT 'session',
  `nick_name` varchar(32) NOT NULL COMMENT '微信昵称',
  `avatar_url` varchar(240) NOT NULL COMMENT '微信头像',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户表' AUTO_INCREMENT=3 ;

--
-- Dumping data for table `mc_user`
--

INSERT INTO `mc_user` (`id`, `openid`, `session`, `nick_name`, `avatar_url`) VALUES
(1, '1x2Bw118Qqwx', '0818Qqw11pab', '', ''),
(2, '1zmvIR1c2IIM', '001c2IIR14BI', 'Colin', 'https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTLTc7AzUyqH2f8VJPmwyno1wnpMRicVLhicI9ymV9cH3vStDHsj1cLHSIgW0UziaMalkMLSXaEL1dibTA/132');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
