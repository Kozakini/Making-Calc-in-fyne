package main
import(
    "fmt"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/theme"
)

type myTheme struct {
    fyne.Theme
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
    if name == theme.SizeNameText {
        return 18 // Default text size
    }
    return m.Theme.Size(name)
}
func main(){
    var lista[] string
    var list[] int
    ap:=app.New()
    win:=ap.NewWindow("calculator")
    ap.Settings().SetTheme(&myTheme{theme.DefaultTheme()})
    wynn:=0
    wys:=widget.NewLabel("")
    je:=widget.NewButton("1",func(){
        elo:=wys.Text+"1"
        wys.Text=elo
        list=append(list,1)
        wys.Refresh()
    })
    dw:=widget.NewButton("2",func(){
        elo:=wys.Text+"2"
        wys.Text=elo
        wys.Refresh()
        list=append(list,2)
    })
    sum:=widget.NewButton("+",func(){
        elo:=wys.Text+"+"
        lista=append(lista,"+")
        wys.Text=elo
        wys.Refresh()
    })
    diff:=widget.NewButton("-",func(){
        elo:=wys.Text+"-"
        wys.Text=elo
        wys.Refresh()
        lista=append(lista,"-")
    })

    tim:=widget.NewButton("*",func(){
        elo:=wys.Text+"*"
        wys.Text=elo
        lista=append(lista,"*")
        wys.Refresh()
    })

    tre:=widget.NewButton("3",func(){
        elo:=wys.Text+"3"
        wys.Text=elo
        list=append(list,3)
        wys.Refresh()
    })
   fo:=widget.NewButton("4",func(){
        elo:=wys.Text+"4"
        wys.Text=elo
        wys.Refresh()
        list=append(list,4)
    })
   fiv:=widget.NewButton("5",func(){
        elo:=wys.Text+"5"
        wys.Text=elo
        wys.Refresh()
        list=append(list,5)
    })

   six:=widget.NewButton("6",func(){
        elo:=wys.Text+"6"
        wys.Text=elo
        wys.Refresh()
        list=append(list,6)
    })
     sev:=widget.NewButton("7",func(){
        elo:=wys.Text+"7"
        wys.Text=elo
        wys.Refresh()
        list=append(list,7)
    })
     eigh:=widget.NewButton("8",func(){
        elo:=wys.Text+"8"
        wys.Text=elo
        wys.Refresh()
        list=append(list,8)
    })
     nine:=widget.NewButton("9",func(){
        elo:=wys.Text+"9"
        wys.Text=elo
        wys.Refresh()
        list=append(list,9)
    })

     zero:=widget.NewButton("0",func(){
        elo:=wys.Text+"0"
        wys.Text=elo
        wys.Refresh()
        list=append(list,0)
    })
    div:=widget.NewButton("/",func(){
            elo:=wys.Text+"/"
            lista=append(lista,"/")
            wys.Text=elo
            wys.Refresh()
    })
    c:=widget.NewButton("C",func(){
        list=[]int{}
        lista=[]string{}
        wys.Text=" "
        wys.Refresh()
    })

    eql:=widget.NewButton("=",func(){
        wynn=list[0]
        for i := 0; i < len(lista); i++{
            if lista[i]=="+"{
                    wynn=wynn+list[i+1]
                }
            if lista[i]=="-"{
                    wynn=wynn-list[i+1]
                }
            if lista[i]=="*"{
                    wynn=wynn*list[i+1]
                }
            if lista[i]=="/"{
                    wynn=wynn/list[i+1]
                }
        }
         wys.Text=fmt.Sprintln(wynn)
         wys.Refresh()
         lista=[]string{}
         list=[]int{}
    })
    e:=container.NewHBox(
        je,
        dw,
        tre,
        sum,
        tim,
    )

    win.SetContent(
        container.NewVBox(
            wys,
            container.NewHBox(
                e,
            ),
            container.NewHBox(
                fo,
                fiv,
                six,
                diff,
                div,
            ),
            container.NewHBox(
                sev,
                eigh,
                nine,
                zero,
                c,
                eql,
                ),
        ),
    )
    win.Resize(fyne.NewSize(600,600))
    win.ShowAndRun()
}


