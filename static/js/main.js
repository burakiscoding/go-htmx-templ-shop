const productAddFormContainer = document.getElementById(
  "product-add-form-container"
);

document.body.addEventListener("htmx:beforeSwap", function (evt) {
  if (evt.detail.xhr.status === 404) {
    // alert the user when a 404 occurs (maybe use a nicer mechanism than alert())
    alert("Error: Could Not Find Resource");
  } else if (evt.detail.xhr.status === 422) {
    // allow 422 responses to swap as we are using this as a signal that
    // a form was submitted with bad data and want to rerender with the
    // errors
    //
    // set isError to false to avoid error logging in console
    evt.detail.shouldSwap = true;
    evt.detail.isError = false;
  } else if (evt.detail.xhr.status === 418) {
    // if the response code 418 (I'm a teapot) is returned, retarget the
    // content of the response to the element with the id `teapot`
    evt.detail.shouldSwap = true;
    evt.detail.target = htmx.find("#teapot");
  }
});

productAddFormContainer.addEventListener("htmx:afterRequest", function (evt) {
  if (evt.detail.xhr.status === 200) {
    const form = document.getElementById("product-add-form");
    if (form) {
      form.reset();
    }
  }
});
