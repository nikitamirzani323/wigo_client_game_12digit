<script>
  import Home from "./pages/Home.svelte"

  export let path_api = "";
  export let path_websocket = "";
  export let version = "";

  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);
  const token_browser = urlParams.get("token");
  
  if (token_browser === null) {
    console.log("TOKEN NOT FOUND");
  } else {
    initTimezone();
  }
  
  let flag_game = false;
  let client_credit = 0;
  let client_ipaddress = "0.0.0.0";
  let client_device = "";
  let client_timezone = "";
  let client_company = "";
  let client_username = "";
  let client_name = "";
  let client_listbet = [];
  let engine_time = 0;
  let engine_time_css = "link-accent";
  let engine_minbet = 0;
  let engine_multiplier_angka = 0;
  let engine_multiplier_redblack = 0;
  let engine_multiplier_line = 0;
  let engine_result = "";
  let engine_status_game_redblackline = "N";
  let engine_status_maintenance = "N";
  let engine_invoice = "Memuat...";
  let engine_status = "LOCK";
  let isModalAlert = false;
  if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
    client_device = "MOBILE";
  } else {
    client_device = "WEBSITE";
  }
  async function initTimezone() {
    const res = await fetch(path_api+"api/healthz");
    if (!res.ok) {
      const message = `An error has occured: ${res.status}`;
      throw new Error(message);
    } else {
      const json = await res.json();
      client_ipaddress = json.real_ip;
      client_timezone = "Asia/Jakarta";
    }
    initapp(token_browser);
  }
  async function initapp(token) {
      const res = await fetch(path_api+"api/checktoken", {
          method: "POST",
          headers: {
              "Content-Type": "application/json",
          },
          body: JSON.stringify({
            client_token: token,
          }),
      });
      const json = await res.json();
      if (json.status === 400) {
      } else if (json.status == 403) {
      } else {
          flag_game = true;
          client_company = json.client_company;
          client_name = json.client_name;
          client_username = json.client_username;
          client_credit = json.client_credit;
          engine_multiplier_angka = json.engine_multiplier_angka;
          engine_multiplier_redblack = json.engine_multiplier_redblack;
          engine_multiplier_line = json.engine_multiplier_line;
          engine_status_game_redblackline = json.engine_status_game_redblackline;
          engine_status_maintenance = json.engine_status_maintenance;
          let record_listbet = json.client_listbet.record;
      
          for (var i = 0; i < record_listbet.length; i++) {
            if(i==0){
              engine_minbet = record_listbet[i]["money_bet"]
            }
            client_listbet = [
                ...client_listbet,
                {
                  money_bet: record_listbet[i]["money_bet"],
                },
            ];
          }
          if(engine_status_maintenance == "Y"){
            isModalAlert = true
            flag_game = false
          }else{
            flag_game = true
            // sse()
            websocket(client_company)
          }
          
      }
  }
  function sse(){
    var source = new EventSource(path_api+"sse");
          source.onmessage = function(event) {
            let text_dasar = event.data;
            let text_replace1 = text_dasar.replace(`"`,"")
            let text_replace2 = text_replace1.replace(`"`,"")
            let text_finalsplit = text_replace2.split("|");
            let data_invoice = text_finalsplit[0];
            let data_time = text_finalsplit[1];
            let data_status = text_finalsplit[2];
            let maintenance_status = text_finalsplit[3];
            let data_result = text_finalsplit[4];

            // console.log(text_dasar)
            if(data_invoice != ""){
              engine_invoice = data_invoice;
            }else{
              engine_invoice = "Memuat...";
            }
            engine_time = data_time;
            if(engine_time < 6){
              engine_time_css = "link-error"
            }else{
              engine_time_css = "link-accent"
            }
            if(data_result != ""){
              engine_result = data_result;
            }else{
              engine_result = "";
            }
            engine_status = data_status;
           
          };
  }
  let conn;
  function websocket(e){
    if (window["WebSocket"]) {
        conn = new WebSocket("wss://"+path_websocket+"/ws/time");
        
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            // appendLog(item);
        };
        conn.onopen = function(evt) {
          // console.log(evt)
          conn.send("nuke")
        }
        conn.onmessage = function (evt) {
            var messages = evt.data;
            let text_replace1 = messages.replace(`"`,'')
            let text_replace2 = text_replace1.replace(`"`,'')
            let text_finalsplit = text_replace2.split("|");
            

            let data_invoice = text_finalsplit[0];
            let data_time = text_finalsplit[1];
            let data_status = text_finalsplit[2];
            let maintenance_status = text_finalsplit[3];
            let data_result = text_finalsplit[4];

            // console.log(text_dasar)
            if(data_invoice != ""){
              engine_invoice = data_invoice;
            }else{
              engine_invoice = "Memuat...";
            }
            engine_time = data_time;
            if(engine_time < 6){
              engine_time_css = "link-error"
            }else{
              engine_time_css = "link-accent"
            }
            if(data_result != ""){
              engine_result = data_result;
            }else{
              engine_result = "";
            }
            engine_status = data_status;
        };
    } else {
        console.log("Your browser does not support WebSockets.")
        // appendLog(item);
    }
  }
  
</script>
{#if engine_status_maintenance == "N"}
  <main class="container mx-auto px-2 mt-5 text-base-content   xl:mt-7 max-w-screen-xl  pb-5 h-fit lg:h-full">
    {#if flag_game}
    <Home 
      {path_api}  
      {engine_time}  
      {engine_time_css}  
      {engine_invoice}  
      {engine_status}  
      {client_listbet}  
      {client_ipaddress}  
      {client_device}  
      {client_company}  
      {client_name}  
      {client_username}  
      {client_credit}  
      {engine_minbet}  
      {engine_multiplier_angka}  
      {engine_multiplier_redblack}  
      {engine_multiplier_line}  
      {engine_result}  
      {engine_status_game_redblackline}  />
    {/if}
  </main>
  <footer class="footer footer-center p-4 text-base-content mt-1 text-center select-none">
    <div class="grid">
      <p class="text-xs lg:text-sm text-center">
        {version}
        <br />
        PowerBy
      </p>
      <img src="https://i.imgur.com/PNSe1ov.png" alt="SDSB TANGKAS" class="w-24 lg:w-28">
    </div>
  </footer>
{/if}

<input type="checkbox" id="my-modal-alert" class="modal-toggle" bind:checked={isModalAlert}>
<div class="modal" on:click|self={()=>isModalAlert = false}>
    <div class="modal-box relative w-11/12 max-w-lg overflow-auto select-none">
        <label for="my-modal-alert" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
        <h3 class="text-xs lg:text-sm font-bold -mt-2">INFO</h3>
        <div class="h-fit overflow-auto  mt-2" >
          <p class="py-4 text-center">
            Website Maintenance
          </p>
        </div>
    </div>
</div>