package xtract

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestProvider_Id(t *testing.T) {

	for _, tt := range []struct {
		in  string
		out string
	}{
		{"https://www.xvideos.com/video17710747/rdt-244", "17710747"},
		{"https://www.xvideos.com/video15689121/abp-386_60a", "15689121"},
		{"https://www.xvideos.com/video29230029/_", "29230029"},
		{"https://www.xvideos.com/video10338658/milf_shizu_amazes_in_raw_solo", "10338658"},
		{"https://www.xvideos.com/video3862306/sweet_kirei_hayakawa_giving_hot_blowjob", "3862306"},

		{"http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=mide00456/", "mide00456"},
		{"http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=mizd00024/?i3_ref=list&i3_ord=2", "mizd00024"},
		{"http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=118npv00007/?i3_ref=list&i3_ord=3", "118npv00007"},
		{"http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=1sdmu00653/?i3_ref=list&i3_ord=4", "1sdmu00653"},
		{"http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=pgd00949/?i3_ref=list&i3_ord=13", "pgd00949"},

		{"https://jp.xhamster.com/movies/3767649/frisky_couple_strip_and_fuck_on_the_beach_hardcore.Html", "3767649"},
		{"https://jp.xhamster.com/videos/she-said-condom-8027212", "8027212"},
		{"https://jp.xhamster.com/videos/taboo-social-statement-preview-by-amedee-vause-8078434", "8078434"},
		{"https://jp.xhamster.com/videos/enko-8077592", "8077592"},

		{"https://www.redtube.com/2338318", "2338318"},
		{"https://redtube.com/143610", "143610"},

		{"https://www.tube8.com/asian/schlong-in-asuka%27s-snatch-%28uncensored-jav%29/32787311/", "asian/schlong-in-asuka%27s-snatch-%28uncensored-jav%29/32787311"},
		{"http://tube8.com/asian/school-girl/185517/", "asian/school-girl/185517"},

		{"https://www.youporn.com/watch/13967749/explosive-vacuum-blowjobs-1-mp4/", "13967749"},
		{"https://www.youporn.com/watch/13951187/serious-cock-sucking-bedroom-porn-with-a-needy-schoolgirl/", "13951187"},

		{"https://jp.pornhub.com/view_video.php?viewkey=ph591db6c3ca15b", "ph591db6c3ca15b"},
		{"https://jp.pornhub.com/view_video.php?viewkey=ph597ae88a7e653", "ph597ae88a7e653"},
		{"https://jp.pornhub.com/view_video.php?viewkey=ph59495aee0e7ad", "ph59495aee0e7ad"},
		{"https://jp.pornhub.com/pornstars", ""},
		{"https://jp.pornhub.com/categories", ""},
	} {
		id := Provider(tt.in).Id()
		assert.Equal(t, tt.out, id)
	}
}

func TestProvider_Embed(t *testing.T) {

	for _, tt := range []struct {
		in  string
		out string
	}{
		{"https://jp.pornhub.com/view_video.php?viewkey=ph591db6c3ca15b", `<iframe src="https://jp.pornhub.com/embed/ph591db6c3ca15b" frameborder="0" Width="560" Height="315" scrolling="no" allowfullscreen></iframe>`},
		{"https://www.xvideos.com/video29230029/_", `<iframe src="https://flashservice.xvideos.com/embedframe/29230029" frameborder=0 Width=510 Height=400 scrolling=no allowfullscreen=allowfullscreen></iframe>`},
		{"http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=mide00456/", `<iframe src="http://www.dmm.co.jp/litevideo/-/part/=/cid=mide00456/size=476_306/" Width="476" Height="306" scrolling="no" frameborder="0" allowfullscreen></iframe>`},
		{"https://jp.xhamster.com/videos/enko-8077592", `<iframe src="https://xhamster.com/xembed.php?video=8077592" Width="510" Height="400" frameborder="0" scrolling="no" allowfullscreen></iframe>`},
		{"https://redtube.com/143610", `<iframe src="https://embed.redtube.com/?id=143610&bgcolor=000000" frameborder="0" Width="560" Height="315" scrolling="no" allowfullscreen></iframe>`},
		{"http://tube8.com/asian/school-girl/185517/", `<iframe src="https://www.tube8.com/embed/asian/school-girl/185517" frameborder="0" Width="608" Height="342" scrolling="no" allowfullscreen="true" webkitallowfullscreen="true" mozallowfullscreen="true" name="t8_embed_video"></iframe>`},
		{"https://www.youporn.com/watch/13951187/serious-cock-sucking-bedroom-porn-with-a-needy-schoolgirl/", `<iframe src='https://www.youporn.com/embed/13951187' frameborder=0 Width='560' Height='315' scrolling=no name='yp_embed_video'></iframe>`},
	} {
		actual := Provider(tt.in).Embed(&Size{}).Html
		assert.Equal(t, tt.out, actual)
	}
}
