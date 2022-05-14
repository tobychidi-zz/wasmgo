const go = new Go();

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
      go.run(result.instance);

      // const data = await featchMovies();

      // renderData(data.results);

      searchApp();
   }
);

// async function featchMovies() {
//    const response = await fetch(
//       "https://moviesdatabase.p.rapidapi.com/titles/search/keyword/love?" +
//          new URLSearchParams({
//             info: "mini_info",
//             limit: "50",
//             page: "2",
//             titleType: "movie",   
//          }),
//       {
//          headers: {
//             "X-RapidAPI-Host": "moviesdatabase.p.rapidapi.com",
//             "X-RapidAPI-Key": "38a48feaa9mshaf0b6e964d158fcp14da0ajsn901265d4d1b0",
//          },
//       }
//    );
//    const data = await response.json();

//    return data;
// }

// function Article(item, target) {
//    const article = document.createElement("article");
//    article.classList.add("column");
//    article.style.maxWidth = "300px";
//    article.style.minWidth = "300px";
//    article.innerHTML = /*html*/ `
//       <div class="card">
//          <div class="card-image">
//             <img src="${
//                item.primaryImage?.url ?? "https://via.placeholder.com/300x400/?text=Movie"
//             }" class="img-responsive" style=" width: 300px; height: 400px; object-fit: cover;">
//       </div>
//       <div class="card-header">
//          <h5 class="card-title">${item.titleText.text}</h5>
//          <p class="card-subtitle text-gray">${item.titleType.text}</p>
//       </div>
//    </div>`;
//    target.appendChild(article);
// }

// function renderData(data) {
//    const target = document.querySelector("#target");
//    target.innerHTML = "";
//    data.forEach((item) => Article(item, target));
// }

function searchApp() {
   const input = document.querySelector("#searchInput");
   input.oninput = (e) => {
      searchNow(e.target.value)
   };
}
