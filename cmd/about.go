package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/teamssix/cf/pkg/cloud"
)

func init() {
	RootCmd.AddCommand(aboutCmd)
}

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "关于作者 (About me)",
	Long:  `关于作者 (About me)`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Print(`
嗨, 我是 TeamsSix，很开心您能找到这儿，您可以在下面的平台中找到并关注我。
Hi, I'm TeamsSix and I'm glad you've found this place. You can find and follow me on the social platforms and links below.

`)
		data := [][]string{
			{"@teamssix", "TeamsSix", "teamssix.com", "github.com/teamssix", "wiki.teamssix.com", "狼组安全团队 @wgpsec"},
		}
		var header = []string{"推特 (Twitter)", "微信公众号", "博客 (Blog)", "Github", "云安全知识库 T Wiki", "所属团队 (Organization)"}
		var td = cloud.TableData{Header: header, Body: data}
		cloud.PrintTable(td, "")
		color.Print(`
如果您使用着感觉还不错，记得给个 Star 哦<gray>（另外 T Wiki 是我自己在维护的云安全知识库，如果您想加入云安全交流群，可以找我，我给你拉进去，在我的公众号里回复“交个朋友”就可以联系到我了。）</>
<bold>感谢您使用我的工具 (Thank you for using my tool)</>
`)
	},
}
