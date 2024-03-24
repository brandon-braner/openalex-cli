CREATE INDEX concepts_ancestors_concept_id_idx ON openalex.concepts_ancestors USING btree (concept_id);


--
-- Name: concepts_related_concepts_concept_id_idx; Type: INDEX; Schema: openalex; Owner: -
--

CREATE INDEX concepts_related_concepts_concept_id_idx ON openalex.concepts_related_concepts USING btree (concept_id);


--
-- Name: concepts_related_concepts_related_concept_id_idx; Type: INDEX; Schema: openalex; Owner: -
--

CREATE INDEX concepts_related_concepts_related_concept_id_idx ON openalex.concepts_related_concepts USING btree (related_concept_id);

--
-- Name: works_primary_locations_work_id_idx; Type: INDEX; Schema: openalex; Owner: -
--

CREATE INDEX works_primary_locations_work_id_idx ON openalex.works_primary_locations USING btree (work_id);


--
-- Name: works_locations_work_id_idx; Type: INDEX; Schema: openalex; Owner: -
--

CREATE INDEX works_locations_work_id_idx ON openalex.works_locations USING btree (work_id);


--
-- Name: works_best_oa_locations_work_id_idx; Type: INDEX; Schema: openalex; Owner: -
--

CREATE INDEX works_best_oa_locations_work_id_idx ON openalex.works_best_oa_locations USING btree (work_id);
