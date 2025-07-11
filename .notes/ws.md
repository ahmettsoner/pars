func executeFunc(cmd *cobra.Command, args []string) error {
	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	workspaceList, err := workspaceService.List()
	if err != nil {
		return fmt.Errorf("Failed to retrieve workspace\n%w", err)
	}

	fmt.Printf("(%d) workspace available\n", len(*workspaceList))

	activeWorkspace, err := workspaceService.GetActiveWorkspace()
	if err != nil {
		return fmt.Errorf("Failed to find Active Workspace\n%w", err)
	}

	selectedWorkspace, err := workspaceService.GetSelectedWorkspace()
	if err != nil {
		return fmt.Errorf("Failed to find Selected Workspace\n%w", err)
	}

	fmt.Println()
	if activeWorkspace != nil && selectedWorkspace != nil {
		if activeWorkspace.Header.Name == selectedWorkspace.Header.Name {
			fmt.Printf("* %v (active & selected)\n", activeWorkspace.Header.Name)
		} else {
			fmt.Printf("* %v (active)\n", activeWorkspace.Header.Name)
			fmt.Printf("%v (selected)\n", selectedWorkspace.Header.Name)
		}
	} else if activeWorkspace != nil {
		fmt.Printf("* %v (active)\n", activeWorkspace.Header.Name)
	} else if selectedWorkspace != nil {
		fmt.Printf("* %v (selected)\n", selectedWorkspace.Header.Name)
	}

	for _, workspace := range *workspaceList {
		if (activeWorkspace == nil || activeWorkspace.Header.Name != workspace.Header.Name) &&
			(selectedWorkspace == nil || selectedWorkspace.Header.Name != workspace.Header.Name) {
			fmt.Println(workspace.Header.Name)
		}
	}

	return nil
}


-----------------



func executeFunc(cmd *cobra.Command, args []string) error {

	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	workspace, err := workspaceService.GetByName(commandOptions.Name)
	if err != nil {
		return fmt.Errorf("Failed to retrieve workspace '%s'\n%w", commandOptions.Name, err)
	}

	if workspace == nil {
		fmt.Println("There are no workspace yet...")
	} else {
		projectService := ioc.Get[application_project_contract.ProjectInterface]()
		projectList, err := projectService.ListByWorkspace(workspace.Specifications.Name)
		if err != nil {
			return fmt.Errorf("Failed to retrieve workspace projects '%s'\n%w", commandOptions.Name, err)
		}

		if commandOptions.PathOnly {
			fmt.Print(workspace.Specifications.Path)
		}

		fmt.Printf("Workspace (%v) has %d project\n", workspace.Header.Name, len(*projectList))
		fmt.Printf("Path : %v \n", workspace.Specifications.Path)

		fmt.Printf("\nProjects:\n")
		if commandOptions.WorkspaceDescribeView.Value == "flat" {
			for _, e := range *projectList {
				name := fmt.Sprintf(" - %v", e.GetFullInformation())
				fmt.Println(name)
			}
		} else if commandOptions.WorkspaceDescribeView.Value == "hierarchical" {
			groups := make(map[string][]string)
			keys := []string{}

			for _, e := range *projectList {
				name := e.GetInformation()
				if !_string.IsEmpty(e.Specifications.GroupObject.Name) {
					groups[e.Specifications.Group] = append(groups[e.Specifications.Group], name)
				} else {
					groups[name] = []string{}
				}
			}

			for key := range groups {
				keys = append(keys, key)
			}
			sort.Strings(keys)

			for _, key := range keys {
				groupItems := groups[key]

				if len(groupItems) > 0 {
					fmt.Printf("%s\n", key)
					for _, value := range groupItems {
						fmt.Printf("  - %s\n", value)
					}
				} else {
					fmt.Printf("- %s\n", key)
				}
			}
		}
	}

	return nil
}