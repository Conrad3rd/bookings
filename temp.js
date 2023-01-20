document.getElementById("colorButton").addEventListener("click", function () {
  let html = `
    <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
      <div class="d-flex flex-row justify-content-evenly" id="reservation-dates-modal">
        <div class="p-2">
          <input disabled required class="form-control" type="text" name="start" id="start" placeholder="Arrival">
        </div>

        <div class="p-2">
          <input disabled required class="form-control" type="text" name="end" id="end" placeholder="Departure">
        </div>
      </div>
    </form>
  `;

  attention.custom({
    title: "Choose your dates",
    msg: html,
  });
})
