import { client } from './baseAxios';

interface ModelCard {
    model: string;
    description: string;
}

interface GetModelsLibraryResponse {
    data: {
        models: ModelCard[];
    };
}

export const getModelsLibrary = async (): Promise<GetModelsLibraryResponse> => {
    const resp = await client.get<GetModelsLibraryResponse>('/models/library');
    return resp.data;
};