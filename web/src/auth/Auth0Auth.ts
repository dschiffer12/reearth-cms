import { useAuth0 } from "@auth0/auth0-react";

import AuthHook from "./AuthHook";

export const errorKey = "reeartherror";

export const useAuth0Auth = (): AuthHook => {
  const {
    user,
    isAuthenticated,
    error,
    isLoading,
    loginWithRedirect,
    logout,
    getAccessTokenSilently,
  } = useAuth0();

  return {
    user,
    isAuthenticated: !!window.REEARTH_E2E_ACCESS_TOKEN || (isAuthenticated && !error),
    isLoading,
    error: error?.message ?? null,
    getAccessToken: () => getAccessTokenSilently(),
    login: () => loginWithRedirect(),
    logout: () =>
      logout({
        returnTo: error
          ? `${window.location.origin}?${errorKey}=${encodeURIComponent(error?.message)}`
          : window.location.origin,
      }),
  };
};
