# Working with Git and GitHub

Git is an essential version control system that helps developers collaborate effectively, track changes, and maintain code quality. However, improper usage can lead to confusion, conflicts, and inefficiencies. The following practices will help maintain a clean and efficient workflow.

---

## 1. Organizing the Git Repository

- Keep the repository well-structured to ensure smooth collaboration.  
- Use a `.gitignore` file to prevent unnecessary files (e.g., build outputs, IDE configs) from being committed.  
- Make each commit focused on a single logical change.

---

## 2. Writing Semantic Commit Messages

- Use **imperative tone** in commit messages.  
  Example:  
  ```bash
  git commit -m "Add user authentication feature"
  ```  
- Keep messages concise, ideally under **50 characters** (short summary).  
- Use the following types for clarity:  

  - **feat**: New feature  
  - **fix**: Bug fix  
  - **refactor**: Code restructuring without changing functionality  
  - **docs**: Documentation updates  
  - **chore**: Maintenance tasks (e.g., dependency updates)  

---

## 3. Branching Strategy: One Branch, One Feature

- **main** → Production-ready code  
- **develop** → Active development  
- **feature/** → For new features  
- **release/** → For preparing releases  
- **hotfix/** → For urgent fixes  

This ensures isolated development and easier collaboration.

---

## 4. Handling Merge Conflicts

Conflicts occur when multiple people edit the same file differently.  

### Preventing Conflicts
- Always pull the latest changes before starting work:  
  ```bash
  git pull origin main
  ```
- Prefer **rebasing** instead of merging for a cleaner history:  
  ```bash
  git rebase main
  ```

### Resolving Conflicts
1. Identify conflicts with:  
   ```bash
   git status
   ```  
2. Manually edit conflicted sections.  
3. Mark as resolved:  
   ```bash
   git add <file>
   ```  
4. Continue rebase or merge:  
   ```bash
   git rebase --continue
   ```

---

## 5. Interactive Rebasing

Before merging your work, use interactive rebasing to clean up commits:  

```bash
git rebase -i HEAD~3
```  

This lets you **squash commits, reword messages, and keep history readable**.

---

## 6. Documenting Changes

- Maintain a detailed **README** file.  
- Add **API documentation** where needed.  
- Keep a **changelog** for tracking updates.  

Clear documentation helps both current team members and new contributors onboard quickly.
