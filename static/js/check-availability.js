document.getElementById("check-availability-button").addEventListener("click", function(){
    //notify("congratulations!!!","success")
    let html = `
    <form id="check-availability" action="" method="" novalidate class="needs-validation" style="overflow: hidden;">
        <div class="row">
            <div class="col" id="reservation-dates-modal" style="height: 300px">
                <div class="row">
                    <div class="col mb-3">
                        <input disabled required class="form-control" type="text" name="start" id="start" placeholder="Arrival" style="width:200px;">
                    </div>
                    <div class="col mb-3">
                        <input disabled required class="form-control" type="text" name="start" id="end" placeholder="Departure" style="width:200px;">
                    </div>

                    <img src="static/images/airplane.gif" alt="">
                </div>
            </div>
        </div>
    </form>
    `
    attention.custom({
        msg: html, 
        title: "Choose your dates",
        callback: function(result){
            console.log("called")

            let form = document.getElementById("check-availability")
            
            let formData = new FormData(form)
            formData.append("csrf_token", "{{.CSRFToken}}")

            fetch('/search-availability-json', {
                method: "post",
                body: formData,
            })
                .then(response => response.json()) //convert to json
                .then(data => { //data = json
                    console.log(data)
                    console.log(data.ok)
                    console.log(data.message)
                })
        }
    })
})