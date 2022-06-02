
Telegram.WebApp.ready();
Telegram.WebApp.MainButton.setParams({
  text_color: '#fff'
});

Telegram.WebApp.onEvent('mainButtonClicked', ()=>{
  
  data = {
    data: {
    name: document.getElementById("name").value,
    url: document.getElementById("url").value,
    method: document.getElementById("method").value,
    body: document.getElementById("body").value,
    interval: document.getElementById("interval").value,
  },
  "button_text":"New monitor"
};
  //window.alert(JSON.stringify(data))
  Telegram.WebApp.sendData("Hello world");
})

updateButton();

function updateButton() {
  Telegram.WebApp.MainButton.setParams({
    is_visible: true,
    text: 'ADD MONITOR',
    color: '#31b545'
  }).onClick(SendData);
}
function toggleBody(){
  document.getElementById("bodyDiv").style = document.getElementById("method").value=='POST'?"display: block":""
}

function SendData(){

    data = {
    name: document.getElementById("name").value,
    url: document.getElementById("url").value,
    method: document.getElementById("method").value,
    body: document.getElementById("body").value,
    interval: document.getElementById("interval").value,
  }

  fetch('https://tools.lavina.uz:8085/addmonitor', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify(data),
}).then(()=>{
  Telegram.WebApp.close() 
})
}

