// Function to make a GET request with Authorization header.
export const doGet = async <T>(url: string): Promise<T> => {
  const headers = new Headers({
    Authorization: `Bearer ${localStorage.getItem('accessToken')}`,
  });

  const requestOptions: RequestInit = {
    method: 'GET',
    headers: headers,
  };

  try {
    const response = await fetch(url, requestOptions);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return (await response.json()) as T;
  } catch (error) {
    console.error('Error:', error);
    throw error; // Re-throw to handle it in the caller function
  }
};

// Function to make a POST request with Authorization header and JSON body.
export const doPost = async <T>(
  url: string,
  data: Record<string, any>,
): Promise<T> => {
  const headers = new Headers({
    Authorization: `Bearer ${localStorage.getItem('accessToken')}`,
    'Content-Type': 'application/json',
  });

  const requestOptions: RequestInit = {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(data), // Convert data to JSON string
  };

  try {
    const response = await fetch(url, requestOptions);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return (await response.json()) as T;
  } catch (error) {
    console.error('Error:', error);
    throw error; // Re-throw to handle it in the caller function
  }
};

export const doDelete = async <T>(url: string): Promise<T> => {
  const headers = new Headers({
    Authorization: `Bearer ${localStorage.getItem('accessToken')}`,
  });

  const requestOptions: RequestInit = {
    method: 'DELETE',
    headers: headers,
  };

  try {
    const response = await fetch(url, requestOptions);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    if (response.status === 204) {
      return {} as T;
    }

    return (await response.json()) as T;
  } catch (error) {
    console.error('Error:', error);
    throw error; // Re-throw to handle it in the caller function
  }
};