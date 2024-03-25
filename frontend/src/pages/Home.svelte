<script>
    import dayjs from "dayjs";
    import utc from "dayjs/plugin/utc";
    import timezone from "dayjs/plugin/timezone";
    import toast, { Toaster } from 'svelte-french-toast';
    
    dayjs.extend(utc);
    dayjs.extend(timezone);
   
    export let path_api = "";
    export let engine_time = 0
    export let engine_time_css = ""
    export let engine_invoice = ""
    export let engine_status = "LOCK"
    export let client_company = ""
    export let client_username = ""
    export let client_timezone = "Asia/Jakarta"
    export let client_name = "";
    export let client_ipaddress = "";
    export let client_device = "";
    export let client_listbet = [];
    export let client_credit = 0;
    export let engine_minbet = 0;
    export let engine_multiplier_angka = 0;
    export let engine_multiplier_redblack = 0;
    export let engine_multiplier_line = 0;
    export let engine_result = "";
    export let engine_status_game_redblackline = "";
    
   
    
    let clockmachine = "";
    
    let flag_btnbuy = true;
    let field_bet = 0
    
    let list_invoice = []
    let list_result = []
    let keranjang = [];
    let flag_listinvoice = true;
    let flag_listresult = false;
    let bet_multiple = []
    let isModalMinBet = false;
    let isModalInfo = false;
    
    let listangka = []
    let redblack = []
    let btn_red_css = "btn btn-error"
    let btn_red_flag = false
    let btn_black_css = "btn"
    let btn_black_flag = false
    let btn_ganjil_css = "btn"
    let btn_ganjil_flag = false
    let btn_genap_css = "btn"
    let btn_genap_flag = false
    let btn_kecil_css = "btn"
    let btn_kecil_flag = false
    let btn_besar_css = "btn"
    let btn_besar_flag = false

    let btn_line1_css = "btn"
    let btn_line1_flag = false
    let btn_line2_css = "btn"
    let btn_line2_flag = false
    let btn_line3_css = "btn"
    let btn_line3_flag = false
    let btn_line4_css = "btn"
    let btn_line4_flag = false
    let btn_line5_css = "btn"
    let btn_line5_flag = false


    function updateClock() {
      let endtime = dayjs().tz(client_timezone).format("DD MMM YYYY | HH:mm:ss");
      clockmachine = endtime;
    }
  
    
    $: {
        setInterval(updateClock, 1000);
        field_bet = engine_minbet
        if(client_username != "" && client_company != ""){
            fetch_invoiceall()
        }
        if(engine_result != ""){
            fetch_invoiceall()
        }
    }
    async function call_bayar() {
        keranjang = [];
        let flag = true;
        let msg_err = ""
      
        let total_bet_multiple = listangka.length
        let total_bayar = parseInt(total_bet_multiple)*parseInt(field_bet)
        let multiplier = 0;
        if(parseInt(engine_time) < 5){
            flag = false
            msg_err = "Timeout"
        }
        if(field_bet == ""){
            flag = false
            msg_err = "Bet wajib diisi"
        }
        if(parseInt(total_bayar) <= 0){
            flag = false
            msg_err = "Nomor dan Bet wajib diisi"
        }
        if(parseInt(field_bet) < parseInt(engine_minbet)){
            flag = false
            msg_err = "Minimal Bet " + engine_minbet
        }
       
        if(parseInt(field_bet) > parseInt(client_credit)){
            flag = false
            msg_err = "Credit tidak cukup "
        }
        if(parseInt(total_bayar) > parseInt(client_credit)){
            flag = false
            msg_err = "Credit tidak cukup "
        }
        // flag = false;
        if(flag){
            flag_btnbuy = false;
            
            for(let i=0;i<total_bet_multiple;i++){
                let tipebet = "ANGKA"
                multiplier = parseFloat(engine_multiplier_angka)
                if(listangka[i] == "RED" || listangka[i] == "BLACK"){
                    tipebet = "REDBLACK"
                    multiplier = parseFloat(engine_multiplier_redblack)
                }
                if(listangka[i] == "GANJIL" || listangka[i] == "GENAP"){
                    tipebet = "REDBLACK"
                    multiplier = parseFloat(engine_multiplier_redblack)
                }
                if(listangka[i] == "KECIL" || listangka[i] == "BESAR"){
                    tipebet = "REDBLACK"
                    multiplier = parseFloat(engine_multiplier_redblack)
                }
                if(listangka[i] == "LINE1" || listangka[i] == "LINE2"){
                    tipebet = "LINE"
                    multiplier = parseFloat(engine_multiplier_line)
                }
                if(listangka[i] == "LINE3" || listangka[i] == "LINE4"){
                    tipebet = "LINE"
                    multiplier = parseFloat(engine_multiplier_line)
                }
                if(listangka[i] == "LINE5"){
                    tipebet = "LINE"
                    multiplier = parseFloat(engine_multiplier_line)
                }
                const data = {
                    ipaddress: client_ipaddress,
                    mobile: "",
                    device: client_device,
                    tipebet: tipebet,
                    nomor: listangka[i],
                    bet: parseInt(field_bet),
                    multiplier: parseFloat(multiplier)
                };
                keranjang = [data, ...keranjang];
            }

            const res = await fetch(path_api+"api/savetransaksi", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    transaksidetail_company: client_company,
                    transaksidetail_idtransaksi: engine_invoice,
                    transaksidetail_username: client_username,
                    transaksidetail_totalbet: parseInt(total_bayar),
                    transaksidetail_listdatabet: keranjang,
                }),
            });
            const json = await res.json();
            if (json.status === 400) {
                flag_btnbuy = true;
            } else if (json.status == 403) {
                alert(json.message);
                flag_btnbuy = true;
            } else {
                client_credit = parseInt(client_credit) - parseInt(total_bayar)
                fetch_invoiceall()
           
                flag_btnbuy = true;
                switch(json.message){
                    case "Success":
                        toast.success(json.message);
                        break;
                    case "Failed":
                        toast.error(msg_err);
                        break;
                    default:
                        toast.error("System Error, Please contact administrator");
                        break;

                }       


                
                call_reset()
            }
        }else{
            
            toast.error(msg_err);
        }
    }
    async function fetch_listresult() {
        flag_listinvoice = false;
        flag_listresult = true;
        list_result = []
        const res = await fetch(path_api+"api/listresult", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                invoice_company: client_company.toLowerCase(),
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
        
        } else if (json.status == 403) {
            alert(json.message);
        } else {
        let record = json.record;
        if (record != null) {
            for (var i = 0; i < record.length; i++) {
                list_result = [
                    ...list_result,
                    {
                        result_invoice: record[i]["result_invoice"],
                        result_date: record[i]["result_date"],
                        result_result: record[i]["result_result"],
                    },
                ];
            }
        }
        }
    }
    async function fetch_invoiceall() {
        flag_listinvoice = true;
        flag_listresult = false;
        list_invoice = []
        const res = await fetch(path_api+"api/listinvoice", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                invoice_company: client_company.toLowerCase(),
                Invoice_username: client_username,
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
        
        } else if (json.status == 403) {
            alert(json.message);
        } else {
        let record = json.record;
        if (record != null) {
            for (var i = 0; i < record.length; i++) {
                let status_css = ""
                
                switch(record[i]["invoiceclient_status"]){
                    case "LOSE":
                        status_css = "badge badge-primary";
                        break;
                    case "WIN":
                        status_css = "badge badge-success";
                        break;
                    case "RUNNING":
                        status_css = "badge badge-warning";
                        break;
                }

               
                list_invoice = [
                    ...list_invoice,
                    {
                        invoiceclient_id: record[i]["invoiceclient_id"],
                        invoiceclient_date: record[i]["invoiceclient_date"],
                        invoiceclient_result: record[i]["invoiceclient_result"],
                        invoiceclient_nomor: record[i]["invoiceclient_nomor"],
                        invoiceclient_tipebet: record[i]["invoiceclient_tipebet"],
                        invoiceclient_bet: record[i]["invoiceclient_bet"],
                        invoiceclient_multiplier: record[i]["invoiceclient_multiplier"],
                        invoiceclient_win: record[i]["invoiceclient_win"],
                        invoiceclient_status: record[i]["invoiceclient_status"],
                        invoiceclient_status_css: status_css,
                    },
                ];
            }
        }
        }
    }
   
    const handleclick_infogame = () => {
        isModalInfo = true
    };
    const handleclick_listminbet = () => {
        isModalMinBet = true
    };
    const handle_minbet = (e) => {
        field_bet = e
        isModalMinBet = false
    };
    
    const call_reset = () => {
        keranjang = [];
        listangka = []
        redblack = []
        field_bet = engine_minbet

        btn_red_css = "btn btn-error"
        btn_red_flag = false
        btn_black_css = "btn"
        btn_black_flag = false
        btn_ganjil_css = "btn"
        btn_ganjil_flag = false
        btn_genap_css = "btn"
        btn_genap_flag = false
        btn_kecil_css = "btn"
        btn_kecil_flag = false
        btn_besar_css = "btn"
        btn_besar_flag = false

        btn_line1_css = "btn"
        btn_line1_flag = false
        btn_line2_css = "btn"
        btn_line2_flag = false
        btn_line3_css = "btn"
        btn_line3_flag = false
        btn_line4_css = "btn"
        btn_line4_flag = false
        btn_line5_css = "btn"
        btn_line5_flag = false

        for(let i=0;i<nomor_master.length;i++){
            nomor[i].nomor_flag = nomor_master[i].nomor_flag
            nomor[i].nomor_css = nomor_master[i].nomor_css
        }
        for(let i=0;i<nomorkecilgenapganjil_master.length;i++){
            nomorkecilgenapganjil[i].nomor_flag = nomorkecilgenapganjil_master[i].nomor_flag
            nomorkecilgenapganjil[i].nomor_css = nomorkecilgenapganjil_master[i].nomor_css
        }
        for(let i=0;i<nomorline_master.length;i++){
            nomorline[i].nomor_flag = nomorline_master[i].nomor_flag
            nomorline[i].nomor_css = nomorline_master[i].nomor_css
        }
    };
    const call_allinvoice = () => {
        fetch_invoiceall()
    };
    const call_listresult = () => {
        fetch_listresult()
    };
    
    const handleclick_redblack = (e) => {
        let objIndex = nomorkecilgenapganjil.findIndex(obj => obj.nomor_id == e);
        let btn_flag = nomorkecilgenapganjil[objIndex].nomor_flag;
        let btn_css = "";
       
        switch(e){
            case "KECIL":
                btn_css = "btn";
                break;
            case "RED":
                btn_css = "btn btn-error";
                break;
            case "BLACK":
                btn_css = "btn";
                break;
            case "GANJIL":
                btn_css = "btn";
                break;
            case "GENAP":
                btn_css = "btn";
                break;
            case "BESAR":
                btn_css = "btn";
                break;
        }
        if(btn_flag == false){
            nomorkecilgenapganjil[objIndex].nomor_flag=true
            nomorkecilgenapganjil[objIndex].nomor_css="btn btn-outline"
            listangka.push(e)
        }else{
            nomorkecilgenapganjil[objIndex].nomor_flag=false
            nomorkecilgenapganjil[objIndex].nomor_css=btn_css
            for (let i = 0; i < listangka.length; i++) { 
                if (listangka[i] === e) { 
                    listangka.splice(i, 1); 
                } 
            }
        }
    };
    const handleclick_line = (e) => {
        let objIndex = nomorline.findIndex(obj => obj.nomor_id == e);
        let btn_flag = nomorline[objIndex].nomor_flag;
        let btn_css = "";
        switch(e){
            case "LINE1":
                btn_css = "btn";
                break;
            case "LINE2":
                btn_css = "btn";
                break;
            case "LINE3":
                btn_css = "btn";
                break;
        }
        if(btn_flag == false){
            nomorline[objIndex].nomor_flag=true
            nomorline[objIndex].nomor_css="btn btn-outline"
            listangka.push(e)
        }else{
            nomorline[objIndex].nomor_flag=false
            nomorline[objIndex].nomor_css=btn_css
            for (let i = 0; i < listangka.length; i++) { 
                if (listangka[i] === e) { 
                    listangka.splice(i, 1); 
                } 
            }
        }
    };
    const handleclick_angka = (e) => {
        let objIndex = nomor.findIndex(obj => obj.nomor_id == e);
        let btn_flag = nomor[objIndex].nomor_flag;
        let btn_css = "";
        switch(e){
            case "00":
                btn_css = "btn btn-error";
                break;
            case "01":
                btn_css = "btn";
                break;
            case "02":
                btn_css = "btn btn-error";
                break;
            case "03":
                btn_css = "btn";
                break;
            case "04":
                btn_css = "btn btn-error";
                break;
            case "05":
                btn_css = "btn";
                break;
            case "06":
                btn_css = "btn btn-error";
                break;
            case "07":
                btn_css = "btn";
                break;
            case "08":
                btn_css = "btn btn-error";
                break;
            case "09":
                btn_css = "btn";
                break;
            case "10":
                btn_css = "btn btn-error";
                break;
            case "11":
                btn_css = "btn";
                break;
        }
        if(btn_flag == false){
            nomor[objIndex].nomor_flag=true
            nomor[objIndex].nomor_css="btn btn-outline"
            listangka.push(e)
        }else{
            nomor[objIndex].nomor_flag=false
            nomor[objIndex].nomor_css=btn_css
            for (let i = 0; i < listangka.length; i++) { 
                if (listangka[i] === e) { 
                    listangka.splice(i, 1); 
                } 
            }
        }
    };
    let nomor_master = [
		{nomor_id: "01", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "02", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "03", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "04", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "05", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "06", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "07", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "08", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "09", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "10", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "11", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "BLACK"},
        {nomor_id: "12", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "RED"},
    ]
    let nomorkecilgenapganjil_master = [
        {nomor_id: "KECIL", nomor_flag:false,nomor_css:"btn"},
		{nomor_id: "GANJIL", nomor_flag:false,nomor_css:"btn"},
		{nomor_id: "BLACK", nomor_flag:false,nomor_css:"btn"},
		{nomor_id: "RED", nomor_flag:false,nomor_css:"btn btn-error"},
		{nomor_id: "GENAP", nomor_flag:false,nomor_css:"btn"},
		{nomor_id: "BESAR", nomor_flag:false,nomor_css:"btn"},
    ]
    let nomorline_master = [
        {nomor_id: "LINE1", nomor_flag:false,nomor_css:"btn"},
        {nomor_id: "LINE2", nomor_flag:false,nomor_css:"btn"},
        {nomor_id: "LINE3", nomor_flag:false,nomor_css:"btn"},
    ]
    let nomor = [
        {nomor_id: "01", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "02", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "03", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "04", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "05", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "BLACK"},
		{nomor_id: "06", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "KECIL", nomor_line: "LINE3", nomor_redblack: "RED"},
		{nomor_id: "07", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE1", nomor_redblack: "BLACK"},
		{nomor_id: "08", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE1", nomor_redblack: "RED"},
		{nomor_id: "09", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE2", nomor_redblack: "BLACK"},
		{nomor_id: "10", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE2", nomor_redblack: "RED"},
		{nomor_id: "11", nomor_flag:false,nomor_css:"btn",nomor_gangen: "GANJIL", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "BLACK"},
        {nomor_id: "12", nomor_flag:false,nomor_css:"btn btn-error",nomor_gangen: "GENAP", nomor_besarkecil: "BESAR", nomor_line: "LINE3", nomor_redblack: "RED"},
    ]
    let nomorkecilgenapganjil = [
        {nomor_id: "KECIL", nomor_flag:false,nomor_css:"btn"},
		{nomor_id: "GANJIL", nomor_flag:false,nomor_css:"btn"},
		{nomor_id: "BLACK", nomor_flag:false,nomor_css:"btn"},
		{nomor_id: "RED", nomor_flag:false,nomor_css:"btn btn-error"},
		{nomor_id: "GENAP", nomor_flag:false,nomor_css:"btn"},
		{nomor_id: "BESAR", nomor_flag:false,nomor_css:"btn"},
    ]
    let nomorline = [
        {nomor_id: "LINE1", nomor_flag:false,nomor_css:"btn"},
        {nomor_id: "LINE2", nomor_flag:false,nomor_css:"btn"},
        {nomor_id: "LINE3", nomor_flag:false,nomor_css:"btn"},
    ]
    function nomorresult(e){
        let css = ""
        for(let i=0;i<nomor_master.length;i++){
            if(e == nomor_master[i].nomor_id){
                css = nomor_master[i].nomor_css
            }
        }
        for(let i=0;i<nomorkecilgenapganjil_master.length;i++){
            if(e == nomorkecilgenapganjil_master[i].nomor_id){
                css = nomorkecilgenapganjil_master[i].nomor_css
            }
        }
        for(let i=0;i<nomorline_master.length;i++){
            if(e == nomorline_master[i].nomor_id){
                css = nomorline_master[i].nomor_css
            }
        }
        return css
    }
    
    function nomorresulttwo(e){
        let css = ""
        for(let i=0;i<nomor_master.length;i++){
            if(e == nomor_master[i].nomor_id){
                css = nomor_master[i].nomor_redblack
            }
        }
        
        if(css == "BLACK"){
            css = "bg-base-200 text-white"
        }else if(css == "RED"){
            css = "bg-error text-black "
        }
        return css
    }
  </script>
 
<section class="glass bg-opacity-60 rounded-lg">
    <section class="flex-col w-full p-2 rounded-md ">
        <section class="flex justify-between w-full">
            <div class="w-full ">
                <center>
                    <img class="w-[150px]" src="https://i.imgur.com/r6BsNCq.png" alt="" srcset="">
                </center>
            </div>
            <div class="w-1/6 lg:w-1/12">
                <button on:click={() => {
                    handleclick_infogame();
                }}  class="btn btn-circle">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                        <path stroke-linecap="round" stroke-linejoin="round" d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z" />
                    </svg>
                </button>
            </div>
        </section>
        <section class="hidden lg:flex justify-between w-full bg-base-100 p-1 rounded-md select-none mt-1">
            <section class="flex-col text-center  font-bold  w-1/2  ">
                <div class="flex-col">
                    <div class="text-lg lg:text-xl">PERIODE</div>
                    <div class="link-accent text-sm lg:text-lg">{engine_invoice}</div>
                </div>
                <div class="flex-col mt-2">
                    <div class="text-lg lg:text-xl">WAKTU</div>
                    <div class="{engine_time_css} text-sm lg:text-lg">{engine_time} S </div>
                </div>
            </section>
            <section class="w-full ">
                <p class="w-full text-[12px] text-right select-none mt-2">
                    Asia/Jakarta <br />
                    {clockmachine}  WIB (+7)<br>
                    {client_name} <br />
                    {client_ipaddress}
                </p>
                <div class="w-full text-[12px] text-right select-none">
                    CREDIT : IDR <span class="link-accent" style="--value:15;">{new Intl.NumberFormat().format(client_credit)}</span>
                </div>
            </section>
        </section>
        <section class="flex-col lg:hidden w-full ">
            <section class="flex justify-between  w-full  bg-base-100 p-2 rounded-md select-none mt-1">
                <div class="flex-col w-full text-center">
                    <div class="text-sm">PERIODE</div>
                    <div class="link-accent text-lg">{engine_invoice}</div>
                </div>
                <div class="flex-col w-full text-center">
                    <div class="text-sm">WAKTU</div>
                    <div class="{engine_time_css} text-lg">{engine_time} S </div>
                </div>
            </section>
            <section class="flex w-full bg-base-100 p-2 rounded-md select-none mt-1">
                <p class="w-full text-xs lg:text-sm text-left select-none">
                    Asia/Jakarta <br />
                    {clockmachine}  WIB (+7)<br>
                    {client_name} <br />
                    {client_ipaddress}<br />
                    CREDIT : IDR <span class="link-accent" style="--value:15;">{new Intl.NumberFormat().format(client_credit)}</span>
                </p>
            </section>
        </section>
        {#if engine_status == "OPEN"}
            <center class="mt-2 bg-base-300 p-2 rounded-2xl text-[11px] lg:text-[15px]">
                Taruhan saya : Pilih beberapa angka untuk bertaruh
            </center>
            <section class="grid grid-cols-1 w-full gap-2 mt-2">
                <div class="h-[270px] lg:h-[230px] w-full overflow-auto">
                    {#if engine_status_game_redblackline == "Y"}
                    <div class="grid grid-cols-3 sm:grid-cols-6 md:grid-cols-6 xl:grid-cols-6 lg:grid-cols-6 gap-1">
                        {#each nomorkecilgenapganjil as rec}
                            <button  on:click={() => {
                                handleclick_redblack(rec.nomor_id);
                            }} class="{rec.nomor_css}">{rec.nomor_id}</button>
                        {/each}
                    </div>
                    <div class="grid grid-cols-3 mt-2  gap-1">
                        {#each nomorline as rec}
                            <button  on:click={() => {
                                handleclick_line(rec.nomor_id);
                            }} class="{rec.nomor_css}">{rec.nomor_id}</button>
                        {/each}
                        
                    </div>
                    {/if}
                    <div class="grid grid-cols-6 mt-2  gap-1 w-full">
                        {#each nomor as rec}
                        <button  on:click={() => {
                                handleclick_angka(rec.nomor_id);
                            }} class="{rec.nomor_css}">{rec.nomor_id}</button>
                        {/each}
                    </div>
                </div>
                <div class="flex-col">
                    <div on:click={() => {
                        handleclick_listminbet();
                    }} class="flex-col w-full h-[90px] justify-center bg-base-300 cursor-pointer rounded-lg">
                        <div class="w-full p-2   text-center ">
                            <div class="uppercase text-xs">Pilih Coin Bet :</div>
                            <div class="text-5xl link-accent">{new Intl.NumberFormat().format(field_bet)}</div>
                        </div>
                    </div>
                </div>
                {#if flag_btnbuy}
                    <div class="grid grid-cols-2 gap-2 w-full">
                        <button on:click={() => {
                                    call_reset();
                            }}  class="btn btn-primary">
                            Reset 
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                            </svg>
                        </button>
                        <button on:click={() => {
                                    call_bayar();
                            }}  class="btn btn-success">
                            Bayar 
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 3h1.386c.51 0 .955.343 1.087.835l.383 1.437M7.5 14.25a3 3 0 0 0-3 3h15.75m-12.75-3h11.218c1.121-2.3 2.1-4.684 2.924-7.138a60.114 60.114 0 0 0-16.536-1.84M7.5 14.25 5.106 5.272M6 20.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Zm12.75 0a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z" />
                            </svg>
                        </button>
                    </div>
                {/if}
            </section>
        {:else}
            <section class="grid grid-cols-1 w-full gap-2 mt-2 h-1/2 text-center ">
                <div class="text-[50px]">
                    RESULT : 
                </div>
                <div class="-mt-5">
                    <span class="rounded-box text-[200px]  font-extrabold {nomorresulttwo(engine_result)}">
                        {engine_result}
                    </span>
                </div>
            </section> 
        {/if}
    </section>
</section>
<section class="flex-col gap-2 mt-4 p-2 glass bg-opacity-60 rounded-md">
    <div class="flex justify-center gap-2">
        <button on:click={() => {
            call_allinvoice();
         }}  class="btn btn-sm">Taruhan Saya</button>
        <button on:click={() => {
            call_listresult();
         }}  class="btn btn-sm">Riwayat</button>
    </div>
    <section class="  mt-4 p-1">
        {#if flag_listinvoice}
        <div class="overflow-auto h-[500px]  scrollbar-thin scrollbar-thumb-green-100">
            <table class="table table-xs  w-full " >
                <thead class="sticky top-0">
                    <tr class="border-none">
                        <th width="5%" class="text-xs text-center align-top">STATUS</th>
                        <th width="5%" class="text-xs text-left align-top">INVOICE</th>
                        <th width="5%" class="text-xs text-center align-top">DATE</th>
                        <th width="10%" class="text-xs text-center align-top">RESULT</th>
                        <th width="*" class="text-xs text-center align-top">NOMOR</th>
                        <th width="10%" class="text-xs text-right align-top">BET</th>
                        <th width="10%" class="text-xs text-right align-top">&nbsp;</th>
                        <th width="10%" class="text-xs text-right align-top">WIN</th>
                    </tr>
                </thead>
                <tbody>
                    {#each list_invoice as rec}
                    <tr class="border-none">
                        <td class="text-xs   text-center whitespace-nowrap align-top">
                            <span class="{rec.invoiceclient_status_css} p-1 text-xs   uppercase   w-20 text-black ">{rec.invoiceclient_status}</span>
                        </td>
                        <td class="text-xs   text-left whitespace-nowrap align-top border-b-transparent">{rec.invoiceclient_id}</td>
                        <td class="text-xs   text-center whitespace-nowrap align-top">{rec.invoiceclient_date}</td>
                        <td class="text-xs   text-center whitespace-nowrap align-top">
                           <span class="{nomorresult(rec.invoiceclient_result)} text-xs  btn-xs">{rec.invoiceclient_result}</span>
                        </td>
                        <td class="text-xs  text-center whitespace-nowrap align-top">
                            <span class="{nomorresult(rec.invoiceclient_nomor)} text-xs  btn-xs">{rec.invoiceclient_nomor}</span>
                        </td>
                        <td class="text-xs  text-right  whitespace-nowrap align-top link-accent {rec.invoice_winlose_css}">{new Intl.NumberFormat().format(rec.invoiceclient_bet)}</td>
                        <td class="text-xs  text-right  whitespace-nowrap align-top link-accent {rec.invoice_winlose_css}">{new Intl.NumberFormat().format(rec.invoiceclient_multiplier)}</td>
                        <td class="text-xs  text-right  whitespace-nowrap align-top link-secondary {rec.invoice_winlose_css}">{new Intl.NumberFormat().format(rec.invoiceclient_win)}</td>
                    </tr>
                    {/each}
                </tbody>
            </table>
        </div>
        {/if}
        {#if flag_listresult}
        <div class="overflow-auto h-[500px]  ">
            <table class="table table-xs w-full " >
                <thead class="sticky top-0">
                    <tr class="border-none">
                        <th width="5%" class="text-xs text-left align-top">INVOICE</th>
                        <th width="5%" class="text-xs text-center align-top">DATE</th>
                        <th width="10%" class="text-xs text-center align-top">RESULT</th>
                    </tr>
                </thead>
                <tbody>
                    {#each list_result as rec}
                    <tr class="border-none">
                        <td class="text-xs lg:text-sm  text-left whitespace-nowrap align-top border-b-transparent">{rec.result_invoice}</td>
                        <td class="text-xs lg:text-sm  text-center whitespace-nowrap align-top">{rec.result_date}</td>
                        <td class="text-xs lg:text-sm  text-center whitespace-nowrap align-top">
                            <span class="{nomorresult(rec.result_result)} text-xs lg:text-sm btn-xs">{rec.result_result}</span>
                        </td>
                    </tr>
                    {/each}
                </tbody>
            </table>
        </div>
        {/if}
    </section>
</section>
<Toaster />

<input type="checkbox" id="my-modal-information" class="modal-toggle" bind:checked={isModalMinBet}>
<div class="modal" on:click|self={()=>isModalMinBet = false}>
    <div class="modal-box relative w-11/12 max-w-lg h-1/2 lg:h-2/3 overflow-auto select-none">
        <label for="my-modal-information" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
        <h3 class="text-xs lg:text-sm font-bold -mt-2">COIN BET</h3>
        <div class="h-fit overflow-auto  mt-2" >
            <div class="grid grid-cols-3 lg:grid-cols-5 mt-5 gap-2 justify-self-center">
                {#each client_listbet as rec}
                <div on:click={() => {
                    handle_minbet(rec.money_bet);
                  }} 
                  class="btn btn-xs lg:btn-sm btn-outline btn-success cursor-pointer">{new Intl.NumberFormat().format(rec.money_bet)}</div>
              {/each}
            </div>
            
        </div>
    </div>
</div>

<input type="checkbox" id="my-modal-infogame" class="modal-toggle" bind:checked={isModalInfo}>
<div class="modal" on:click|self={()=>isModalMinBet = false}>
    <div class="modal-box relative w-11/12 max-w-xl h-1/2 lg:h-2/3 overflow-auto select-none">
        <label for="my-modal-infogame" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
        <h3 class="text-xs lg:text-sm font-bold -mt-2">INFO</h3>
        <div class="h-fit overflow-auto  mt-2 w-full" >
            <table class="table table-xs w-full ">
                <thead>
                    <tr>
                        <th width="80%">PENJELASAN</th>
                        <th width="*">HADIAH</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>ANGKA (01 - 12)</td>
                        <td>{engine_multiplier_angka} x</td>
                    </tr>
                    <tr>
                        <td>BESAR/KECIL,GENAP/GANJIL,RED/BLACK</td>
                        <td>{engine_multiplier_redblack} x</td>
                    </tr>
                    <tr>
                        <td>LINE 1,2,3</td>
                        <td>{engine_multiplier_line} x</td>
                    </tr>
                </tbody>
            </table>
            <table class="table table-xs w-full mt-2">
                <thead>
                    <tr>
                        <th width="80%">PENJELASAN</th>
                        <th width="*">NOMOR</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>KECIL</td>
                        <td>01-06</td>
                    </tr>
                    <tr>
                        <td>BESAR</td>
                        <td>07-12</td>
                    </tr>
                    <tr>
                        <td>Line 1</td>
                        <td>01,02,07,08</td>
                    </tr>
                    <tr>
                        <td>Line 2</td>
                        <td>03,04,09,10</td>
                    </tr>
                    <tr>
                        <td>Line 3</td>
                        <td>05,06,11,12</td>
                    </tr>
                </tbody>
            </table>

            
            <p class="text-[12px] mt-2">
                <b class="uppercase font-bold">Cara Bermain :</b> <br />
                Pilih angka 01 - 12 <br />
                Nomor akan diundi setelah waktu 0 Second,  <br />
                jika nomor anda kena, maka anda akan mendapatkan: modal + (modal * hadiah)
                <br /><br />
                Contoh :<br />
                KASUS 1:<br />
                Anda memasang nomor 10, dengan bet 1,000<br />
                keluaran adalah nomor 10<br />
                jadi bet anda menang dari hasil Angka : 10, Line 3, Besar, Genap, Red <br />
                pembayarannya adalah : modal + (modal x 5)<br />
                1,000 + (1,000 x 10) = 11,000<br />
                anda akan mendapatkan 11,000
                <br /><br />
                KASUS 2:<br />
                Anda memasang nomor 12, dengan bet 1,000<br />
                keluaran adalah nomor 00<br />
                jadi bet anda kalah 
                <br /><br />
                KASUS 3:<br />
                Anda memasang nomor GENAP, dengan bet 1,000<br />
                keluaran adalah nomor 02<br />
                jadi bet anda menang dari hasil Angka : 02, Line 2, Genap, Kecil, Red <br />
                pembayarannya adalah : modal + (modal x 0.95)<br />
                1,000 + (1,000 x 0.95) = 1,950<br />
                anda akan mendapatkan 1,950
        </div>
    </div>
</div>