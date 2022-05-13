document.onload(()=>{
    window.Telegram.WebApp.MainButton.setParams({
        text: "Add monitor"
    })
    window.Telegram.WebApp.MainButton.show()
    
    window.Telegram.WebApp.expand()
    a=1
    
    document.getElementById('body1').value = JSON.stringify(window.Telegram.WebApp)
    
    alert("Hello")

})