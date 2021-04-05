import { useState, useEffect } from "react";
import { useAuthState } from "./useAuthContext";
const ROOT_URL = "/ui/api";

const useFetch = (url, options) => {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const { token } = useAuthState();

  useEffect(() => {
    const init = async () => {
      try {
        const response = await fetch(`${ROOT_URL}/${url}`, {
          method: "GET",
          headers: new Headers({
            Authorization: `Bearer ${token}`,
          }),
        });

        if (response.ok) {
          const json = await response.json();
          setData(json);
        } else if (response.status === 401) {
          //TODO: fix this hack
          localStorage.removeItem("token");
          localStorage.removeItem("currentUser");
          window.location = "/";
        } else {
          throw response;
        }
      } catch (e) {
        console.error("fetch failed: ", e);
        setError(e);
      } finally {
        setLoading(false);
      }
    };

    init();
  }, [url, token]); // rerun when...

  return { data, error, loading };
};

export default useFetch;