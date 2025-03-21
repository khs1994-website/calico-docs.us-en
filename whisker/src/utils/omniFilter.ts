import {
    OmniFilterOption as ListOmniFilterOption,
    OperatorType,
} from '@/libs/tigera/ui-components/components/common/OmniFilter/types';
import {
    ApiFilterResponse,
    FlowsFilterKeys,
    FlowsFilter,
    FlowsFilterQuery,
    QueryPage,
} from '@/types/api';

export enum FilterKey {
    policy = 'policy',
    source_name = 'source_name',
    source_namespace = 'source_namespace',
    dest_name = 'dest_name',
    dest_namespace = 'dest_namespace',
    protocol = 'protocol',
    port = 'port',
    action = 'action',
}

export type FilterProperty = {
    filterHintsKey: FlowsFilterKeys;
    transformToFilterHintRequest: (
        filters: string[],
    ) => FlowsFilterQuery[] | undefined | string[];
};

export const ListOmniFilterKeys: Omit<
    typeof FilterKey,
    'action' | 'protocol' | 'port'
> = {
    [FilterKey.policy]: FilterKey.policy,
    [FilterKey.source_namespace]: FilterKey.source_namespace,
    [FilterKey.source_name]: FilterKey.source_name,
    [FilterKey.dest_namespace]: FilterKey.dest_namespace,
    [FilterKey.dest_name]: FilterKey.dest_name,
} as const;

export type ListOmniFilterParam = keyof typeof ListOmniFilterKeys;

export const OmniFilterKeys: Omit<typeof FilterKey, 'action'> = {
    ...ListOmniFilterKeys,
    [FilterKey.port]: FilterKey.port,
    [FilterKey.protocol]: FilterKey.protocol,
} as const;

export type OmniFilterParam = keyof typeof OmniFilterKeys;

export const CustomOmniFilterKeys: Pick<typeof FilterKey, 'port'> = {
    [FilterKey.port]: FilterKey.port,
} as const;

export type CustomOmniFilterParam = keyof typeof CustomOmniFilterKeys;

const handleEmptyFilters = (filters: any[]) =>
    filters.length ? filters : undefined;

const transformToListFilter = (
    filters: string[] = [],
): FlowsFilterQuery[] | undefined =>
    handleEmptyFilters(
        filters.map((value) => ({
            type: 'Exact',
            value,
        })),
    );

export const transformToFlowsFilterQuery = (
    omniFilterValues: Record<FilterKey, string[]>,
    listFilterId?: ListOmniFilterParam,
    searchInput?: string,
) => {
    const filterHintsQuery: FlowsFilter = Object.keys(omniFilterValues).reduce(
        (acc, filterKey) => {
            const filterId = filterKey as FilterKey;
            const key = FilterProperties[filterId].filterHintsKey;

            return listFilterId === filterId
                ? acc
                : {
                      ...acc,
                      [key]: FilterProperties[
                          filterId
                      ].transformToFilterHintRequest(
                          omniFilterValues[filterId],
                      ),
                  };
        },
        {},
    );

    if (listFilterId && searchInput) {
        const key = OmniFilterProperties[listFilterId].filterHintsKey;
        filterHintsQuery[key] = [
            {
                type: 'Fuzzy',
                value: searchInput,
            },
        ];
    }

    return Object.keys(filterHintsQuery).length
        ? JSON.stringify(filterHintsQuery)
        : '';
};

export type FilterHintType =
    | 'SourceName'
    | 'DestName'
    | 'DestNamespace'
    | 'SourceNamespace'
    | 'PolicyTier';

export const FilterHintTypes: Record<ListOmniFilterParam, FilterHintType> = {
    [ListOmniFilterKeys.dest_name]: 'DestName',
    [ListOmniFilterKeys.dest_namespace]: 'DestNamespace',
    [ListOmniFilterKeys.source_name]: 'SourceName',
    [ListOmniFilterKeys.source_namespace]: 'SourceNamespace',
    [ListOmniFilterKeys.policy]: 'PolicyTier',
};

export type OmniFilterPropertiesType = Record<
    OmniFilterParam,
    {
        selectOptions?: ListOmniFilterOption[];
        defaultOperatorType?: OperatorType;
        label: string;
        limit?: number;
        filterHintsKey: Exclude<FlowsFilterKeys, 'actions'>;
        transformToFilterHintRequest: (
            filters: string[],
        ) => FlowsFilterQuery[] | undefined;
    }
>;

const requestPageSize = 20;

export const OmniFilterProperties: OmniFilterPropertiesType = {
    policy: {
        label: 'Policy',
        limit: requestPageSize,
        filterHintsKey: 'policies',
        transformToFilterHintRequest: transformToListFilter,
    },
    source_namespace: {
        label: 'Source Namespace',
        limit: requestPageSize,
        filterHintsKey: 'source_namespaces',
        transformToFilterHintRequest: transformToListFilter,
    },
    dest_namespace: {
        label: 'Dest Namespace',
        limit: requestPageSize,
        filterHintsKey: 'dest_namespaces',
        transformToFilterHintRequest: transformToListFilter,
    },
    source_name: {
        label: 'Source',
        limit: requestPageSize,
        filterHintsKey: 'source_names',
        transformToFilterHintRequest: transformToListFilter,
    },
    dest_name: {
        label: 'Destination',
        limit: requestPageSize,
        filterHintsKey: 'dest_names',
        transformToFilterHintRequest: transformToListFilter,
    },
    port: {
        label: 'Port',
        filterHintsKey: 'dest_ports',
        transformToFilterHintRequest: (values: string[]) =>
            handleEmptyFilters(
                values
                    .map(Number)
                    .filter(Boolean)
                    .map((value) => ({ type: 'Exact', value })),
            ),
    },
    protocol: {
        label: 'Protocol',
        filterHintsKey: 'protocols',
        transformToFilterHintRequest: transformToListFilter,
    },
};

export const FilterProperties: Record<FilterKey, FilterProperty> = {
    ...OmniFilterProperties,
    action: {
        filterHintsKey: 'actions',
        transformToFilterHintRequest: handleEmptyFilters,
    },
};

export type ListOmniFiltersData = Record<
    ListOmniFilterParam,
    ListOmniFilterData
>;

export type ListOmniFilterData = {
    filters: ListOmniFilterOption[] | null;
    isLoading: boolean;
    total?: number;
};

export type SelectedOmniFilterData = Partial<ListOmniFiltersData>;

export type SelectedOmniFilters = Partial<Record<OmniFilterParam, string[]>>;

export type SelectedOmniFilterOptions = Record<
    ListOmniFilterParam,
    ListOmniFilterOption[]
>;

export const transformToQueryPage = (
    { items, total }: ApiFilterResponse,
    page: number,
): QueryPage => ({
    items: items.map(({ value }) => ({ label: value, value })),
    total,
    currentPage: page,
    nextPage: page + 1,
});
