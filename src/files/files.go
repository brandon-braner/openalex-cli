package files

func GetValidFilenames() []string {
	validFileNames := []string{
		"authors",
		"authors_counts_by_year",
		"authors_ids",
		"concepts",
		"concepts_ancestors",
		"concepts_counts_by_year",
		"concepts_ids",
		"concepts_related_concepts",
		"institutions",
		"institutions_associated_institutions",
		"institutions_counts_by_year",
		"institutions_geo",
		"institutions_ids",
		"publishers",
		"publishers_counts_by_year",
		"publishers_ids",
		"sources",
		"sources_counts_by_year",
		"sources_ids",
		"topics",
		"works",
		"works_authorships",
		"works_best_oa_locations",
		"works_biblio",
		"works_concepts",
		"works_ids",
		"works_locations",
		"works_mesh",
		"works_open_access",
		"works_primary_locations",
		"works_referenced_works",
		"works_related_works",
		"works_topics"}

	return validFileNames
}
