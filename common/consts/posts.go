package consts

// swagger:enum ArticleType
// ENUM(
// Text = 0,
// Image = 1,
// Audio = 2,
// Video = 3,
// )
type ArticleType int64

func (a ArticleType) Cn() string {
	switch a {
	case ArticleTypeText:
		return "图文"
	case ArticleTypeImage:
		return "图片"
	case ArticleTypeAudio:
		return "音频"
	case ArticleTypeVideo:
		return "视频"
	default:
		return "未知"
	}
}

// swagger:enum ArticleFormat
// ENUM(
// Html = 0,
// Markdown = 1,
// )
type ArticleFormat int64

func (a ArticleFormat) Cn() string {
	switch a {
	case ArticleFormatHtml:
		return "HTML"
	case ArticleFormatMarkdown:
		return "Markdown"
	default:
		return "未知"
	}
}

// swagger:enum PriceType
// ENUM(
// Content = 0,
// Attachment = 1,
// )
type ArticlePriceType int64

func (a ArticlePriceType) Cn() string {
	switch a {
	case ArticlePriceTypeContent:
		return "内容"
	case ArticlePriceTypeAttachment:
		return "附件"
	default:
		return "未知"
	}
}

// swagger:enum VideoProvider
// ENUM(url, bilibili, tencent, youku, youtube, iqiyi, sohu, xigua, douyin, weishi, weibo)
type VideoProvider string

// swagger:enum AudioProvider
// ENUM(url, ximalaya)
type AudioProvider string
