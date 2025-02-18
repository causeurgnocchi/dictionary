import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch, url, params }) => {
  console.log("load");
  return {
    results: await fetch(
      `http://[::1]:8080/api/vocabularies/${params.search}`
    ).then((response) => response.json()),
  };
};
