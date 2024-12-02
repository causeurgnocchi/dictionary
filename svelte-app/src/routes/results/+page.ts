import type { PageLoad } from './$types';

export const load: PageLoad = async ({ url }) => {
  return {
      vocabularies: await fetch(`http://[::1]:8080/api/vocabularies/${url.searchParams.get('search')}`)
      .then(response => response.json())
  }
};
