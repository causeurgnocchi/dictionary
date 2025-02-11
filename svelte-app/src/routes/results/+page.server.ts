import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, url }) => {
  if (url.searchParams.get('search') === '') redirect(302, '/');
  
  return {
    vocabularies: await fetch(`http://[::1]:8080/api/vocabularies/${url.searchParams.get('search')}`).then(response => response.json())
  }
};
