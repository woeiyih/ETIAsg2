async function displayAPIData() {
  //get API data
  const response = await fetch("https://jsonplaceholder.typicode.com/posts");
  data = await response.json();

  //generate HTML code
  const tableData = data
    .map(function (value) {
      return `<tr>
            <td>${value.StudentName}</td>
            <td>${value.StudentID}</td>
            <td>${value.DOB}</td>
            <td>${value.PhoneNumber}</td>
            <td>${value.Address}</td>
        </tr>`;
    })
    .join("");

  //set tableBody to new HTML code
  const tableBody = document.querySelector("#tableBody");
  tableBody.innerHTML = tableData;
}

displayAPIData();
