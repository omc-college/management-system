export interface Endpoint {
  id: number;
  name: string;
  path: string;
  method: string;
}

export interface FeatureEntry {
  id: number;
  name: string;
  description: string;
  endpoints: Endpoint[];
}

export interface Role {
  id: number;
  name: string;
  entries: FeatureEntry[];
}
