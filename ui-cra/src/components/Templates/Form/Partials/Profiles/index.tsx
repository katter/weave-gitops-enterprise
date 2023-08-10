import { Checkbox, TextField } from '@material-ui/core';
import { Autocomplete } from '@material-ui/lab';
import { DataTable, Flex } from '@weaveworks/weave-gitops';
import _ from 'lodash';
import React, { Dispatch, FC, useEffect, useState } from 'react';
import styled from 'styled-components';
import {
  ClusterNamespacedName,
  RepositoryRef,
} from '../../../../../cluster-services/cluster_services.pb';
import { ProfilesIndex, UpdatedProfile } from '../../../../../types/custom';
import { Loader } from '../../../../Loader';
import ProfilesListItem from './ProfileListItem';
import { CheckBoxOutlineBlank, CheckBox } from '@material-ui/icons';

const icon = <CheckBoxOutlineBlank fontSize="small" />;
const checkedIcon = <CheckBox fontSize="small" />;

const ProfilesWrapper = styled.div`
  padding-bottom: ${({ theme }) => theme.spacing.xl};
  .table-wrapper {
    max-height: 500px;
    overflow: scroll;
  }
  table {
    thead {
      th:first-of-type {
        padding: ${({ theme }) => theme.spacing.xs}
          ${({ theme }) => theme.spacing.base};
      }
      h2 {
        line-height: 1;
      }
    }
    td:first-of-type {
      text-overflow: clip;
      width: 25px;
      padding-left: ${({ theme }) => theme.spacing.base};
    }
    a {
      color: ${({ theme }) => theme.colors.primary};
    }
    .profile-details {
      justify-content: space-evenly;
    }
    .MuiFormControl-root {
      min-width: 150px;
      margin-right: 24px;
    }
  }
`;

const ProfileDetailsLabelRenderer = () => (
  <Flex className="profile-details">
    <h2>Version</h2>
    <h2>Namespace</h2>
    <h2>Yaml</h2>
  </Flex>
);

const Profiles: FC<{
  cluster?: ClusterNamespacedName;
  context?: string;
  updatedProfiles: ProfilesIndex;
  setUpdatedProfiles: Dispatch<React.SetStateAction<ProfilesIndex>>;
  isLoading: boolean;
  isProfilesEnabled?: string;
  helmRepos: RepositoryRef[];
}> = ({
  context,
  cluster,
  updatedProfiles,
  setUpdatedProfiles,
  isLoading,
  isProfilesEnabled = 'true',
  helmRepos,
}) => {
  const [selectedHelmRepositories, setSelectedHelmRepositories] = useState<
    RepositoryRef[]
  >([]);

  // We need to check the updated profiles and see if there have been any selections made (callbackState).
  // If there have been,  we need to find those helm repositories and set them as selectedHelmRepositories

  useEffect(() => setSelectedHelmRepositories(helmRepos), [helmRepos]);

  const handleIndividualClick = (
    event: React.ChangeEvent<HTMLInputElement>,
    name: string,
  ) => {
    setUpdatedProfiles(sp => ({
      ...sp,
      [name]: {
        ...sp[name],
        selected: event.target.checked,
      },
    }));
  };

  const handleSelectAllClick = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUpdatedProfiles(sp =>
      _.mapValues(sp, p => ({
        ...p,
        selected: event.target.checked || p.required,
      })),
    );
  };

  const updatedProfilesList = _.sortBy(Object.values(updatedProfiles), [
    'name',
  ]);

  const numSelected = updatedProfilesList.filter(up => up.selected).length;
  const rowCount = updatedProfilesList.length || 0;

  // Showing helm repositories in autocomplete as name:namespace as there could be multiple repositories with same name in different namespaces
  const nameNamespaceHelmRepos = helmRepos.map(
    hr => `${hr.name}:${hr.namespace}`,
  );

  const getSelectedHelmRepos = (selectedHelmRepos: string[]) =>
    selectedHelmRepos.map((selectedHelmRepo: string) => {
      const [name, namespace] = selectedHelmRepo.split(':');
      return helmRepos.find(
        hr => hr.name === name && hr.namespace === namespace,
      ) as RepositoryRef;
    });

  // TO DO: Check Gitlab - do we need to add the selected repos to the state as we do with the profiles?
  // TO DO: Check the form in EDIT mode once we're able to create resources in this new format

  return (
    <ProfilesWrapper>
      <h2>{context === 'app' ? 'Helm Releases' : 'Profiles'}</h2>
      <Autocomplete
        multiple
        id="helmrepositories-select"
        options={nameNamespaceHelmRepos.sort()}
        disableCloseOnSelect
        getOptionLabel={option => option as string}
        onChange={(event, selectedNameNamespaceHelmRepos: string[]) =>
          setSelectedHelmRepositories(
            getSelectedHelmRepos(selectedNameNamespaceHelmRepos),
          )
        }
        renderOption={(option: string, { selected }) => (
          <li>
            <Checkbox
              color="primary"
              icon={icon}
              checkedIcon={checkedIcon}
              style={{ marginRight: 8 }}
              checked={selected}
            />
            {option}
          </li>
        )}
        renderInput={params => (
          <TextField
            {...params}
            label="HelmRepositories"
            placeholder="Helm Repositories"
          />
        )}
      />
      {isLoading && <Loader />}
      {!isLoading && (
        <DataTable
          className="profiles-table table-wrapper"
          rows={updatedProfilesList.filter(
            up =>
              selectedHelmRepositories?.find(
                hr =>
                  hr.name === up.repoName && hr.namespace === up.repoNamespace,
              ) !== undefined,
          )}
          fields={[
            {
              label: 'checkbox',
              labelRenderer: () => (
                <Checkbox
                  onChange={handleSelectAllClick}
                  checked={rowCount > 0 && numSelected === rowCount}
                  indeterminate={numSelected > 0 && numSelected < rowCount}
                  color="primary"
                />
              ),
              value: (profile: UpdatedProfile) => (
                <Checkbox
                  onChange={event => handleIndividualClick(event, profile.name)}
                  checked={Boolean(updatedProfiles[profile.name]?.selected)}
                  disabled={profile.required}
                  color={profile.required ? undefined : 'primary'}
                />
              ),
              maxWidth: 25,
            },
            {
              label: 'Name',
              value: (p: UpdatedProfile) => (
                <span data-profile-name={p.name}>{p.name}</span>
              ),
              sortValue: ({ name }) => name,
              maxWidth: 220,
            },
            ...(context !== 'app'
              ? [
                  {
                    label: 'Layer',
                    value: (p: UpdatedProfile) =>
                      p.layer ? (
                        <div className="profile-layer">
                          <span>{p.layer}</span>
                        </div>
                      ) : null,
                  },
                ]
              : []),
            {
              label: 'Repository',
              value: (p: UpdatedProfile) => (
                <span data-profile-repository={p.repoName}>{p.repoName}</span>
              ),
              sortValue: ({ repoName }) => repoName,
              maxWidth: 220,
            },
            {
              label: 'Version',
              labelRenderer: () => <ProfileDetailsLabelRenderer />,
              value: (p: UpdatedProfile) => (
                <ProfilesListItem
                  className="profile-details"
                  cluster={cluster}
                  context={context}
                  profile={p}
                  setUpdatedProfiles={setUpdatedProfiles}
                  helmRepo={
                    helmRepos.find(
                      hr => hr.name === p.repoName,
                    ) as RepositoryRef
                  }
                />
              ),
            },
          ]}
          hideSearchAndFilters={true}
        />
      )}
    </ProfilesWrapper>
  );
};

export default Profiles;
