import type { PageLoad } from './$types';

export const load: PageLoad = ({ params }) => {
  return {
    post: {
      vocabularies: fetch(`http://[::1]:8080/api/vocabularies/${params.}`)
      .then(response => response.json())
    }
  }
};
