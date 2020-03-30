package downloadurl_test

import "AnimeunityCLI/packages/commonresources"

var (
	testCases = []Test{
		{
			commonresources.AnimePageStruct{"200", "https://animeunity.it/anime.php?id=200", "", nil, false,2000},
			commonresources.AnimePageStruct{
				"200",
				"https://animeunity.it/anime.php?id=200",
				"High School DxD ",
				[]string{
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_12_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_11_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_10_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_09_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_08_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_07_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_06_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_05_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_04_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_03_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_02_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_01_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_Ep_01_SUB_ITA.mp4",
				},
				false,
				2000,
			},
		},
		{
			commonresources.AnimePageStruct{"666", "https://animeunity.it/anime.php?id=666", "", nil, true,2000},
			commonresources.AnimePageStruct{
				"666",
				"https://animeunity.it/anime.php?id=666",
				"High School DxD OVA ",
				[]string{
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_OAV_02_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_OAV_01_SUB_ITA.mp4",
					"https://www.animeunityserver21.cloud/DDL/Anime/HighSchoolDxD/HighSchoolDxD_OAV_01_SUB_ITA.mp4",
				},
				true,
				2000,
			},
		},
		{
			commonresources.AnimePageStruct{"203", "https://animeunity.it/anime.php?id=203", "", nil, false,2000},
			commonresources.AnimePageStruct{
				"203",
				"https://animeunity.it/anime.php?id=203",
				"High School DxD New ",
				[]string{
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_12_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_11_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_10_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_09_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_08_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_07_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_06_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_05_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_04_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_03_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_02_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_01_SUB_ITA.mp4",
					"https://www.animeunityserver96.cloud/DDL/Anime/HighSchoolDxDNew/HighSchoolDxDNew_Ep_01_SUB_ITA.mp4",
				},
				false,
				2000,
			},
		},
	}
)
