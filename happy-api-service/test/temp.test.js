async function tempFetch() {
  const request = await fetch('http://127.0.0.1:55010/api/products');
  const responseJson = await request.json();

  responseJson.data.forEach((product) => {
    console.info(product);
  });
}

tempFetch();
