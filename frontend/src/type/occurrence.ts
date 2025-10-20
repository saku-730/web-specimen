// src/types/occurrence.ts

import type { Time } from 'time'; // time.Timeに対応する型をインポート (必要に応じて調整)

// エクスポートして他のファイルから使えるようにするのだ
export interface ClassificationDetail {
  classification_id: number;
  species?: string;
  genus?: string;
  family?: string;
  order?: string;
  class?: string;
  phylum?: string;
  kingdom?: string;
  others?: string;
}

export interface ObservationDetail {
  observation_id: number;
  observation_user_id: number;
  observation_user: string;
  observation_method_id: number;
  observation_method_name: string;
  page_id?: number;
  behavior?: string;
  observed_at: string; // または Date 型
}

export interface SpecimenDetail {
  specimen_id: number;
  specimen_user_id: number;
  specimen_user: string;
  specimen_methods_id: number;
  specimen_methods_common: string;
  created_at: string; // または Date 型
  page_id?: number;
  institution_id: number;
  institution_code: string;
  collection_id?: string;
}

export interface IdentificationDetail {
  identification_id: number;
  identification_user_id: number;
  identification_user: string;
  identified_at: string; // または Date 型
  source_info?: string;
}

export interface AttachmentDetail {
  attachment_id: number;
  file_path: string;
  file_name?: string;
  note?: string;
}

export interface OccurrenceDetailResponse {
  user_id: number;
  user_name: string;
  project_id: number;
  project_name: string;
  individual_id?: number;
  lifestage?: string;
  sex?: string;
  body_length?: string;
  created_at: string; // または Date 型
  language_id?: number;
  latitude?: number;
  longitude?: number;
  place_name?: string;
  note?: string;
  classification?: ClassificationDetail;
  observation: ObservationDetail[];
  specimen: SpecimenDetail[];
  identification: IdentificationDetail[];
  attachments: AttachmentDetail[];
}

// 注意: Goの time.Time に対応する型は、通常 string (ISO 8601) または Date です。
// 必要に応じて `import type { Time } from 'time';` のようなインポートを追加・調整してください。
